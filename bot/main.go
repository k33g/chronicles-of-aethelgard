package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ollama/ollama/api"

	"github.com/k33g/chronicles-of-aethelgard/ui"
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

	errEnv := godotenv.Load("./data/.env")

	ollamaUrl := os.Getenv("OLLAMA_HOST")
	model := os.Getenv("LLM")

	client, errCli := api.ClientFromEnvironment()

	systemInstructionsFile, errInst := os.ReadFile("./data/instructions.md")
	systemInstructions := string(systemInstructionsFile)

	// Configuration
	configFile, errConf := os.ReadFile("./data/settings.json")
	var config map[string]interface{}
	errJsonConf := json.Unmarshal(configFile, &config)

	errorsList := errors.Join(errEnv, errCli, errInst, errConf, errJsonConf)
	if errorsList != nil {
		log.Fatal("😡:", errorsList)
	}

	ui.Println("#ffc0c5", "🌍", ollamaUrl, "📕", model)
	ui.Println("#FFFF00", "📝 config:", config)

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

		userContent := data["user"]

		// Prompt construction
		messages := []api.Message{
			{Role: "system", Content: systemInstructions},
			//{Role: "user", Content: userContent},
		}

		// Add memory
		messages = append(messages, memory...)
		// Add the new question
		messages = append(messages, api.Message{Role: "user", Content: userContent})

		req := &api.ChatRequest{
			Model:    model,
			Messages: messages,
			Options:  config,
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

		// Save the conversation in memory
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
