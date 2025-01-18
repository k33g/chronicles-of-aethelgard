package services

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Request struct {
	Question string `json:"question"`
}

func SpeakWithGrym(question string) (string, error) {
	return SpeakWithBot(os.Getenv("BOT_GRYM_SERVICE_HOST"), question)
}

func SpeakWithElvira(question string) (string, error) {
	return SpeakWithBot(os.Getenv("BOT_ELVIRA_SERVICE_HOST"), question)
}

func SpeakWithEthan(question string) (string, error) {
	return SpeakWithBot(os.Getenv("BOT_ETHAN_SERVICE_HOST"), question)
}

func SpeakWithBot(botServiceHost, question string) (string, error) {
	botServiceEndPoint := "/api/chat"
	data := Request{
		Question: question,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error marshalling data: %w", err)
	}

	fmt.Printf("Sending data to the AI Bot: %s on %s\n\n", string(jsonData), botServiceHost)

	req, err := http.NewRequest("POST", botServiceHost+botServiceEndPoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Configure client with timeouts and keep-alive
	client := &http.Client{
		Transport: &http.Transport{
			DisableCompression: true,
			ForceAttemptHTTP2:  true,
			MaxIdleConns:       100,
			IdleConnTimeout:    90 * time.Second,
		},
		Timeout: 0, // No timeout for streaming
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Use a larger buffer and scanner for more reliable reading
	reader := bufio.NewReaderSize(resp.Body, 4096) // 4KB buffer
	var answer strings.Builder

	for {
		chunk, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				// Add the last chunk if it exists
				if len(chunk) > 0 {
					answer.Write(chunk)
					fmt.Print(string(chunk))
					os.Stdout.Sync()
				}
				break
			}
			return "", fmt.Errorf("error reading response: %w", err)
		}

		answer.Write(chunk)
		fmt.Print(string(chunk))
		os.Stdout.Sync()
	}

	return answer.String(), nil
}

func _speakWithBot(botServiceHost, question string) (string, error) {
	botServiceEndPoint := "/api/chat"
	// Create a Request
	data := Request{
		Question: question,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("😡 Error marshalling data:", err)
		return "", err
	}
	fmt.Printf("Sending data to the AI Bot: %s on %s\n\n", string(jsonData), botServiceHost)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", botServiceHost+botServiceEndPoint, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("😡 Error creating request:", err)
		return "", err
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
		fmt.Println("😡 Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response in streaming mode
	// Create a buffer to read the response in chunks
	buffer := make([]byte, 32) // Read 32 bytes at a time
	answer := ""
	for {
		n, err := resp.Body.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("😡 Error reading response:", err)
			return "", err
		}
		if n > 0 {
			fmt.Print(string(buffer[:n]))
			answer += string(buffer[:n])
			os.Stdout.Sync()
		}
	}
	fmt.Println()
	return answer, nil
}
