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

	generationInstructions := `Your job is to generate a description of a room in a fantasy setting using the name given by the user.
	The output must be in markdown format, with the name of the room as the title and then the description of the room:
	<Expected Output>
	# Name of the room

	Description of the room
	</Expected Output>

	Ensure the description is fantasy-themed.
	`

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/room/generate/name", func(response http.ResponseWriter, request *http.Request) {
		// add a flusher
		flusher, ok := response.(http.Flusher)
		if !ok {
			response.Write([]byte("😡 Error: expected http.ResponseWriter to be an http.Flusher"))
		}

		userContent := `
		Generate a short random room name for a miedeval dungeon in a D&D game.
		Generate only one unique name, no description is needed.
		<Expected Output>
		The name of the room
		</Expected Output>
		`
		// Prompt construction
		messages := []api.Message{
			{Role: "system", Content: systemInstructions},
			{Role: "user", Content: userContent},
		}

		stream := true
		//noStream  := false

		// Configuration
		options := map[string]interface{}{
			"temperature":   2.0,
			"repeat_last_n": 2,
			//"top_k":         10,
			//"top_p":         0.5,
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

		// Run completion
		err := client.Chat(ctx, req, respFunc)

		if err != nil {
			log.Fatal("😡:", err)
		}
	})

	mux.HandleFunc("POST /api/room/generate/description", func(response http.ResponseWriter, request *http.Request) {

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
		roomName := data["room_name"]
		userContent := "Generate a short room description for the name '" + roomName + "' with the above instructions."

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
