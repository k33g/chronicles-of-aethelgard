package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ollama/ollama/api"
)

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
	// TODO: test if variables are empty

	fmt.Println("🌍", ollamaUrl, "📕", model)

	client, errCli := api.ClientFromEnvironment()
	if errCli != nil {
		log.Fatal("😡:", errCli)
	}

	personality, err := os.ReadFile("./personality.md")
	if err != nil {
		log.Fatal("😡:", err)
	}

	// Define a tool
	escape := map[string]any{
		"type": "function",
		"function": map[string]any{
			"name":        "escape",
			"description": "escape of the place thanks to the magic words",
			"parameters": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"first": map[string]any{
						"type":        "string",
						"description": "The first magic word",
					},
					"second": map[string]any{
						"type":        "string",
						"description": "The second magic word",
					},
					"third": map[string]any{
						"type":        "string",
						"description": "The third magic word",
					},
				},
				"required": []string{"first", "second", "third"},
			},
		},
	}
	tools := []any{escape}
	// transform tools to json
	jsonTools, _ := json.MarshalIndent(tools, "", "  ")

	fmt.Println("🛠️", string(jsonTools))

	// Transform the tools to Ollama format
	var toolsList api.Tools
	jsonErr := json.Unmarshal(jsonTools, &toolsList)
	if jsonErr != nil {
		log.Fatalln("😡", jsonErr)
	}

	systemInstructions := `You are a Sphinx, your name is Abul-Hol, the "Father of Terror",
	expert at interpreting and answering questions based on provided sources.
	Use the below PERSONALITY content to answer user questions. 
	Be verbose and speak like an Egyptian Sphinx!`

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
			{Role: "system", Content: systemInstructions},
			{Role: "system", Content: string(personality)},
			{Role: "user", Content: userContent},
		}

		stream := true
		noStream  := false

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
		if err != nil {
			log.Fatal("😡:", err)
		}
		// first try to add a specific question to exit the place/game
		// detect if the user want to trigger an action (from the user question only)
		// New prompt construction
		messages = []api.Message{
			{Role: "user", Content: userContent},
		}
		req = &api.ChatRequest{
			Model: model, // The model must support the tools
			Messages: messages,
			Options: map[string]interface{}{
				"temperature":   0.0,
				"repeat_last_n": 2,
			},
			Tools:  toolsList,
			Stream: &noStream,
		}
		err = client.Chat(ctx, req, func(resp api.ChatResponse) error {
			//fmt.Println("🖐️", resp.Message.ToolCalls)
			/*
			for _, toolCall := range resp.Message.ToolCalls {
				fmt.Println(toolCall.Function.Name, toolCall.Function.Arguments)
			}
			*/
			if len(resp.Message.ToolCalls) == 0 {
				return nil
			}

			toolCall := resp.Message.ToolCalls[0]
			fmt.Println(toolCall.Function.Name, toolCall.Function.Arguments)
			

			jsonToolCall, _ := json.MarshalIndent(toolCall, "", "  ")
			response.Write([]byte("\n🛠️ :\n" + string(jsonToolCall)))
			//response.Write(jsonToolCall)
			flusher.Flush()

			// Then make the call to the tool

	
			return nil
		})
	
		if err != nil {
			log.Fatalln("😡", err)
		}

	})

	var errListening error
	log.Println("🌍 http server is listening on: " + httpPort)
	errListening = http.ListenAndServe(":"+httpPort, mux)

	log.Fatal(errListening)

}
