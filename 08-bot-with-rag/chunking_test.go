package main

import (
	"08-bot-with-rag/rag"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/ollama/ollama/api"
)

func TestGenerateChunk(t *testing.T) {
	ctx := context.Background()

	ollamaUrl := os.Getenv("OLLAMA_HOST")
	embeddingsModel := os.Getenv("EMBEDDINGS_LLM")

	fmt.Println("🌍", ollamaUrl, "📦", embeddingsModel)

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

	content, err := os.ReadFile("./character-sheet-" + characterSheetId + ".md")
	if err != nil {
		log.Fatal("😡:", err)
	}

	vectorStore := []rag.VectorRecord{}

	chunks := rag.SplitText(string(content), "<!-- SPLIT -->")

	// Create embeddings from documents and save them in the store
	for idx, chunk := range chunks {
		fmt.Println("📝 Creating embedding nb:", idx)
		fmt.Println("📝 Chunk:", chunk)

		embedding, _ := rag.GetEmbeddingFromChunk(ctx, client, embeddingsModel, chunk)

		// Save the embedding in the vector store
		record := rag.VectorRecord{
			Prompt:    chunk,
			Embedding: embedding,
		}
		vectorStore = append(vectorStore, record)
	}

	// Marshal the store to JSON
	storeJSON, err := json.MarshalIndent(vectorStore, "", "  ")
	if err != nil {
		log.Fatal("Failed to marshal store to JSON:", err)
	}

	// Write the JSON to a file
	storeFile := "store-" + characterSheetId + ".json"
	err = os.WriteFile(storeFile, storeJSON, 0644)
	if err != nil {
		log.Fatal("Failed to write store to file:", err)
	}

	fmt.Println("✅ Store persisted to", storeFile)

}
