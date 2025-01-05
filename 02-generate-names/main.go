package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ollama/ollama/api"
)

func main() {

	ctx := context.Background()
	//errEnv := godotenv.Load()
	//if errEnv != nil {
	//	log.Fatal("😡:", errEnv)
	//}

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

	generationInstructions := `Generate a random name for a role-playing game character for a given kind (species/race). The output should be in JSON format, with the keys 'name' and 'kind'. Ensure the name is fantasy-themed.
		**Expected Output:** 
		Generate a JSON object with the keys 'name' and 'kind'. For example:
		{
		"name": "John Doe",
		"kind": "Human"
		}
	`

	//userContent := "Give a name for a Dwarf."
	userContent := "Give a name for an Elf."
	//userContent := "Give a name for a Human."

	// Prompt construction
	messages := []api.Message{
		{Role: "system", Content: systemInstructions},
		{Role: "system", Content: generationInstructions},
		{Role: "user", Content: userContent},
	}

	stream := true
	//noStream  := false

	req := &api.ChatRequest{
		Model:    model,
		Messages: messages,
		Options: map[string]interface{}{
			"temperature":   0.0,
			"repeat_last_n": 2,
			"top_k":         10,
			"top_p":         0.5,
		},
		Format:    json.RawMessage(`"json"`),
		KeepAlive: &api.Duration{Duration: 1 * time.Minute},
		Stream:    &stream,
	}

	jsonResult := ""
	respFunc := func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		jsonResult += resp.Message.Content
		return nil
	}

	// Start the chat completion
	errChat := client.Chat(ctx, req, respFunc)
	if errChat != nil {
		log.Fatal("😡:", errChat)
	}

	errJson := os.WriteFile("./character.json", []byte(jsonResult), 0644)
	if errJson != nil {
		log.Fatal("😡:", errJson)
	}

	fmt.Println("\n🟦")

}
