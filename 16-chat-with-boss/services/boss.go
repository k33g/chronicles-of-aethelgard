package services

import "os"

func SpeakWithSphinx(question string) (string, error) {
	return speakWithBot(os.Getenv("BOT_SPHINX_SERVICE_HOST"), question)
}
