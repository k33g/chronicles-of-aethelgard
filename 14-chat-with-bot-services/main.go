package main

import (
	"14-chat-with-bot-services/services"
)

func main() {
	services.SpeakWithGrym("What is your name?")
	services.SpeakWithGrym("Remember my name, I'm Philippe")
	services.SpeakWithGrym("What is my name?")

	services.SpeakWithElvira("What is your name?")

}
