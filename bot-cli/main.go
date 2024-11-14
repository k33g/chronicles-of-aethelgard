package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Request struct {
	User string `json:"user"`
}

func main() {
	// Load the environment variables
	err := godotenv.Load(filepath.Join("data", ".env"))
	if err != nil {
		log.Fatal("😡 Error loading .env file", err)
	}

	botHost := os.Getenv("BOT_HOST")
	if botHost == "" {
		log.Fatal("😡 BOT_HOST is not set")
	}

	// Create a Request
	data := Request{
		User: "why the sky is blue?",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("😡 Error marshalling data:", err)
	}

	fmt.Printf("Sending data to the AI Bot: %s on %s\n\n", string(jsonData), botHost)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", botHost+"/api/chat", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("😡 Error creating request:", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP client
	client := &http.Client{
		Transport: &http.Transport{
			DisableCompression: true,
		},
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("😡 Error sending request:", err)
	}
	defer resp.Body.Close()

	// Read the response in streaming mode
	// Create a buffer to read the response in chunks
	buffer := make([]byte, 32) // Read 32 bytes at a time
	for {
		n, err := resp.Body.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("😡 Error reading response:", err)
		}
		if n > 0 {
			fmt.Print(string(buffer[:n]))
			os.Stdout.Sync()
		}
	}
	fmt.Println()
}
