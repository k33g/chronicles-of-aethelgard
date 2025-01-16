package main

import (
	"16-chat-with-boss/services"
	"16-chat-with-boss/tools"
	"encoding/json"
	"fmt"
)

type Function struct {
	Name      string    `json:"name"`
	Arguments Arguments `json:"arguments"`
}

type Arguments struct {
	First  string `json:"first"`
	Second string `json:"second"`
	Third  string `json:"third"`
}

type ToolCall struct {
	Function Function `json:"function"`
}

func main() {

	for {
		input := tools.Input("Question (type 'quit' ou 'exit' to quit) : ")

		if input == "quit" || input == "exit" {
			fmt.Println("👋 Bye!")
			break
		}

		sphinxAnswer, _ := services.SpeakWithSphinx(input)
		//services.SpeakWithSphinx(input)
		fmt.Println()

		// Test the Sphinx answer
		var toolCall ToolCall
		err := json.Unmarshal([]byte(sphinxAnswer), &toolCall)
		if err != nil {
			//fmt.Println("😡: Error unmarshalling Sphinx answer")
			// if error it's not a tool call
		} else {
			if toolCall.Function.Arguments.First == "yellow" && toolCall.Function.Arguments.Second == "black" && toolCall.Function.Arguments.Third == "green" {
				fmt.Println()
				fmt.Println("🎉: You escaped!")
				break
			} else {
				fmt.Println()
				fmt.Println("😡: You are still trapped!")
			}
		}

	}
}
