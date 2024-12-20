package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
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

func main() {

	ctx := context.Background()
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("😡:", errEnv)
	}

	ollamaUrl := os.Getenv("OLLAMA_HOST")
	model := os.Getenv("LLM")

	fmt.Println("🌍", ollamaUrl, "📕", model)

	client, errCli := api.ClientFromEnvironment()
	if errCli != nil {
		log.Fatal("😡:", errCli)
	}

	systemInstructions, err := os.ReadFile("instructions.md")
	if err != nil {
		log.Fatal("😡:", err)
	}

	generationInstructions, err := os.ReadFile("steps.md")
	if err != nil {
		log.Fatal("😡:", err)
	}

	// Get the character
	character, errChar := GetCharacter()
	if errChar != nil {
		log.Fatal("😡:", errChar)
	}

	fmt.Println("🧙‍♂️", character.Name, "🧝‍♂️", character.Kind)

	userContent := fmt.Sprintf("Create a %s with this name:%s", character.Kind, character.Name)

	// Prompt construction
	messages := []api.Message{
		{Role: "system", Content: string(systemInstructions)},
		{Role: "system", Content: string(generationInstructions)},
		{Role: "user", Content: userContent},
	}

	stream := true
	//noStream  := false

	req := &api.ChatRequest{
		Model:    model,
		Messages: messages,
		Options: map[string]interface{}{
			//"temperature":   0.0,
			"temperature":   0.8,
			"repeat_last_n": 2,
			"top_k":         10,
			"top_p":         0.5,
		},
		//Format:    "json",
		KeepAlive: &api.Duration{Duration: 1 * time.Minute},
		Stream:    &stream,
	}

	mdResult := ""
	respFunc := func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		mdResult += resp.Message.Content
		return nil
	}

	// Start the chat completion
	errChat := client.Chat(ctx, req, respFunc)
	if errChat != nil {
		log.Fatal("😡:", errChat)
	}

	// Character sheet
	characterSheetId := strings.ToLower(strings.ReplaceAll(character.Name, " ", "-"))

	log.Printf("Attempting to write file: ./character-sheet-%s.md", characterSheetId)

	errWriteFile := os.WriteFile("./character-sheet-"+characterSheetId+".md", []byte("# CHARACTER SHEET\n\n"+mdResult), 0644)
	if errWriteFile != nil {
		log.Fatal("😡:", errChat)
	}

	fmt.Println("📝", characterSheetId, "saved.")

}
