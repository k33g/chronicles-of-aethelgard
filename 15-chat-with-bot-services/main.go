package main

import (
	"15-chat-with-bot-services/services"
	"15-chat-with-bot-services/tools"
	"fmt"
)

func main() {

	for {
		input := tools.Input("Question (type 'quit' ou 'exit' to quit) : ")

		if input == "quit" || input == "exit" {
			fmt.Println("👋 Bye!")
			break
		}

		services.SpeakWithElvira(input)
	}

}
