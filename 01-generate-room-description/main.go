package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ollama/ollama/api"
)

func main() {

	ctx := context.Background()

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

	userContent := "Generate a room description for the name 'The Forgotten Library' with the above instructions."
	/*
		The Grand Entrance
		The Hall of Whispers
		The Chamber of Echoes
		The Forgotten Library
		The Dark Cave
	*/

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
			"temperature":   0.8,
			"repeat_last_n": 2,
			"top_k":         10,
			"top_p":         0.5,
		},
		KeepAlive: &api.Duration{Duration: 1 * time.Minute},
		Stream:    &stream,
	}

	respFunc := func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		return nil
	}

	// Start the chat completion
	errChat := client.Chat(ctx, req, respFunc)
	if errChat != nil {
		log.Fatal("😡:", errChat)
	}

	fmt.Println("\n🟦")
}
