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

	fmt.Println("🌍", ollamaUrl, "📕", model)

	client, errCli := api.ClientFromEnvironment()
	if errCli != nil {
		log.Fatal("😡:", errCli)
	}

	systemInstructions := `# IDENTITY and PURPOSE
		You are an expert NPC generator for games like D&D 5th edition. 
		You have freedom to be creative to get the best possible output.
	`

	generationInstructions := `👋 c'est par ici !
	`

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/monster/generate/description", func(response http.ResponseWriter, request *http.Request) {

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

		//userContent := data["user"]
		monsterName := data["monster_name"]
		userContent := "Generate a short monster description for the name '" + monsterName + "' with the above instructions."

		// Prompt construction
		messages := []api.Message{
			{Role: "system", Content: systemInstructions},
			{Role: "system", Content: generationInstructions},
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
