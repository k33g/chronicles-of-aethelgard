package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ollama/ollama/api"
)

type Character struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

func GetCharacter() (Character, error) {
	// Read the JSON file
	file, errRead := os.ReadFile("./character.json")
	if errRead != nil {
		return Character{}, errRead
	}

	// Unmarshal the JSON data into a struct
	var character Character
	errUnmarshall := json.Unmarshal(file, &character)
	if errUnmarshall != nil {
		return Character{}, errUnmarshall
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

	context, err := os.ReadFile("./character-sheet-" + characterSheetId + ".md")
	if err != nil {
		log.Fatal("😡:", err)
	}

	systemContentTpl := `You are a %s, your name is %s,
	expert at interpreting and answering questions based on provided sources.
	Using only the provided context, answer the user's question 
	to the best of your ability using only the resources provided. 
	Be verbose!`

	systemInstructions := fmt.Sprintf(systemContentTpl, character.Kind, character.Name)

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

		// Prompt construction
		messages := []api.Message{
			{Role: "system", Content: "CONTEXT: " + string(context)},
			{Role: "system", Content: systemInstructions},
			{Role: "user", Content: userContent},
		}

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

		// TODO: add an option to stop the completion

		/*
			_, err = completion.ChatStream(ollamaUrl, query,
				func(answer llm.Answer) error {
					log.Println("📝:", answer.Message.Content)
					response.Write([]byte(answer.Message.Content))

					flusher.Flush()
					if !shouldIStopTheCompletion {
						return nil
					} else {
						return errors.New("🚫 Cancelling request")
					}
				})

		*/

		err = client.Chat(ctx, req, respFunc)

		if err != nil {
			log.Fatal("😡:", err)
		}

	})

	var errListening error
	log.Println("🌍 http server is listening on: " + httpPort)
	errListening = http.ListenAndServe(":"+httpPort, mux)

	log.Fatal(errListening)

}
