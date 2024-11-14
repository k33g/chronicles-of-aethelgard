package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/k33g/chronicles-of-aethelgard/ui"

	"github.com/joho/godotenv"
	"github.com/ollama/ollama/api"
)

func main() {

	brain := ""
	if len(os.Args) < 2 {
		// default brain
		brain = "v5"
	} else {
		brain = os.Args[1]
	}

	ctx := context.Background()
	errEnv := godotenv.Load(fmt.Sprintf("./data/brain-%s/.env", brain))

	ollamaUrl := os.Getenv("OLLAMA_HOST")
	model := os.Getenv("LLM")

	ui.Println("#ffc0c5", "🌍", ollamaUrl, "📕", model)

	client, errCli := api.ClientFromEnvironment()

	// Configuration
	configFile, errConf := os.ReadFile(fmt.Sprintf("./data/brain-%s/settings.json", brain))
	var config map[string]interface{}
	errJsonConf := json.Unmarshal(configFile, &config)

	errorsList := errors.Join(errEnv, errCli, errConf, errJsonConf)
	if errorsList != nil {
		log.Fatal("😡:", errorsList)
	}

	ui.Println("#FFFF00", "📝 config:", config)

	// +++ Define tools +++
	parameters := struct {
		Type       string   `json:"type"`
		Required   []string `json:"required"`
		Properties map[string]struct {
			Type        string   `json:"type"`
			Description string   `json:"description"`
			Enum        []string `json:"enum,omitempty"`
		} `json:"properties"`
	}{
		Type:     "object",
		Required: []string{"a", "b"},
		Properties: map[string]struct {
			Type        string   `json:"type"`
			Description string   `json:"description"`
			Enum        []string `json:"enum,omitempty"`
		}{
			"a": {
				Type:        "number",
				Description: "first operand",
			},
			"b": {
				Type:        "number",
				Description: "second operand",
			},
		},
	}

	tools := []api.Tool{
		{
			Type: "function",
			Function: api.ToolFunction{
				Name:        "addNumbers",
				Description: "Make an addition of the two given numbers",
				Parameters:  parameters,
			},
		},
		{
			Type: "function",
			Function: api.ToolFunction{
				Name:        "multiplyNumbers",
				Description: "Make a multiplication of the two given numbers",
				Parameters:  parameters,
			},
		},
	}


	for {
		question, _ := ui.Input("#008000", fmt.Sprintf("🤖 [%s] 🧠 (%s) ask me something> ", model, brain))

		if question == "bye" {
			break
		}

		// Prompt construction
		messages := []api.Message{
			{Role: "user", Content: question},
		}

		/*
		{Role: "user", Content: `add 2 and 40`},
		{Role: "user", Content: `multiply 2 and 21`},
		*/


		req := &api.ChatRequest{
			Model: model,
			Tools: tools,
			Format:   "json",
			Stream:   new(bool), // set streaming to false
			Messages: messages,
			Options:  config,
		}

		// Start the counter goroutine
		done := make(chan struct{})
		go func() {
			counter := 0
			for {
				select {
				case <-done:
					return
				default:
					counter++
					fmt.Printf("\r⏳ Computing... %d seconds", counter)
					time.Sleep(1 * time.Second)
				}
			}
		}()

		answer := ""

		respFunc := func(resp api.ChatResponse) error {
			if answer == "" {
				fmt.Println(" ✅")
				fmt.Println()
				close(done)
			}
			// TODO: get the first tool call
			fmt.Print(resp.Message.ToolCalls)

			return nil
		}

		err := client.Chat(ctx, req, respFunc)

		if err != nil {
			log.Fatal("😡:", err)
		}
		fmt.Println()
		fmt.Println()

	}

	/*

	*/
}
