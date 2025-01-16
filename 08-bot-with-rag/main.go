package main

import (
	"08-bot-with-rag/rag"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ollama/ollama/api"
)

type Character struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

func GetCharacter() (Character, error) {
	var character Character
	character.Name = os.Getenv("CHARACTER_NAME")
	character.Kind = os.Getenv("CHARACTER_KIND")

	if character.Name == "" || character.Kind == "" {
		return character, fmt.Errorf("😡: character name or kind not set")
	}

	return character, nil
}

/*
GetBytesBody returns the body of an HTTP request as a []byte.
  - It takes a pointer to an http.Request as a parameter.
  - It returns a []byte.
*/
func GetBytesBody(request *http.Request) []byte {
	body := make([]byte, request.ContentLength)
	request.Body.Read(body)
	return body
}

func main() {
	ctx := context.Background()

	var httpPort = os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	ollamaUrl := os.Getenv("OLLAMA_HOST")
	model := os.Getenv("LLM")
	embeddingsModel := os.Getenv("EMBEDDINGS_LLM")

	fmt.Println("🌍", ollamaUrl, "📕", model)

	client, errCli := api.ClientFromEnvironment()
	if errCli != nil {
		log.Fatal("😡:", errCli)
	}

	// Get the character
	character, errChar := GetCharacter()
	if errChar != nil {
		log.Fatal("😡:", errChar)
	}

	fmt.Println("🧙‍♂️", character.Name, "🧝‍♂️", character.Kind)

	characterSheetId := strings.ToLower(strings.ReplaceAll(character.Name, " ", "-"))

	vectorStore := []rag.VectorRecord{}
	// Unmarshal the store from a JSON file if it exists
	storeFile := "store-" + characterSheetId + ".json"
	file, err := os.ReadFile(storeFile)
	if err != nil {
		log.Fatal("😡 Failed to read store file:", err)
	}
	if err := json.Unmarshal(file, &vectorStore); err != nil {
		log.Fatal("😡 Failed to unmarshal store:", err)
	}

	/*
		context, err := os.ReadFile("./character-sheet-" + characterSheetId + ".md")
		if err != nil {
			log.Fatal("😡:", err)
		}
	*/

	systemContentTpl := `You are a %s, your name is %s,
	expert at interpreting and answering questions based on provided sources.
	Using only the provided context, answer the user's question 
	to the best of your ability using only the resources provided. 
	Be verbose!`

	systemInstructions := fmt.Sprintf(systemContentTpl, character.Kind, character.Name)

	// 🧠 Memory
	memory := []api.Message{
		{Role: "system", Content: "CONVERSATION MEMORY:"},
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/chat", func(response http.ResponseWriter, request *http.Request) {

		// add a flusher
		flusher, ok := response.(http.Flusher)
		if !ok {
			response.Write([]byte("😡 Error: expected http.ResponseWriter to be an http.Flusher"))
		}
		body := GetBytesBody(request)
		// unmarshal the json data
		var data map[string]string

		err := json.Unmarshal(body, &data)
		if err != nil {
			response.Write([]byte("😡 Error: " + err.Error()))
		}

		userContent := data["question"]

		// 🧠 Get the context from the similarities
		embeddingFromQuestion, _ := rag.GetEmbeddingFromChunk(ctx, client, embeddingsModel, userContent)

		// Search similarites between the question and the vectors of the store
		// 1- calculate the cosine similarity between the question and each vector in the store
		similarities := []rag.Similarity{}

		for _, vector := range vectorStore {
			cosineSimilarity, err := rag.CosineSimilarity(embeddingFromQuestion, vector.Embedding)
			if err != nil {
				log.Fatalln("😡", err)
			}

			// append to similarities
			similarities = append(similarities, rag.Similarity{
				Prompt:           vector.Prompt,
				CosineSimilarity: cosineSimilarity,
			})
		}

		// Select the 2 most similar chunks
		// retrieve in similarities the 5 records with the highest cosine similarity
		// sort the similarities
		sort.Slice(similarities, func(i, j int) bool {
			return similarities[i].CosineSimilarity > similarities[j].CosineSimilarity
		})

		// get the first 2 records
		top5Similarities := similarities[:5]

		fmt.Println("🔍 Top similarities:")
		for _, similarity := range top5Similarities {
			fmt.Println("🔍 Prompt:", similarity.Prompt)
			fmt.Println("🔍 Cosine similarity:", similarity.CosineSimilarity)
			fmt.Println("--------------------------------------------------")
		}

		// Create a new context with the top 5 chunks
		newContext := ""
		for _, similarity := range top5Similarities {
			newContext += similarity.Prompt
		}

		// Prompt construction
		messages := []api.Message{
			{Role: "system", Content: "CONTEXT: " + string(newContext)},
			{Role: "system", Content: systemInstructions},
			//{Role: "user", Content: userContent},
		}

		// 🧠 Add memory
		messages = append(messages, memory...)
		// Add the new question
		messages = append(messages, api.Message{Role: "user", Content: userContent})

		stream := true
		//noStream  := false

		// Configuration
		options := map[string]interface{}{
			"temperature":   0.8,
			"repeat_last_n": 2,
			"top_k":         10,
			"top_p":         0.5,
		}

		req := &api.ChatRequest{
			Model:     model,
			Messages:  messages,
			Options:   options,
			KeepAlive: &api.Duration{Duration: 1 * time.Minute},
			Stream:    &stream,
		}

		answer := ""
		respFunc := func(resp api.ChatResponse) error {

			response.Write([]byte(resp.Message.Content))

			fmt.Print(resp.Message.Content)
			answer += resp.Message.Content

			flusher.Flush()

			return nil
		}

		err = client.Chat(ctx, req, respFunc)

		// 🧠 Save the conversation in memory
		memory = append(
			memory,
			api.Message{Role: "user", Content: userContent},
			api.Message{Role: "assistant", Content: answer},
		)

		if err != nil {
			log.Fatal("😡:", err)
		}

	})

	var errListening error
	log.Println("🌍 http server is listening on: " + httpPort)
	errListening = http.ListenAndServe(":"+httpPort, mux)

	log.Fatal(errListening)

}
