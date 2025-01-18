package services

import "os"

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


func SpeakWithSphinx(question string) (string, error) {
	return SpeakWithBot(os.Getenv("BOT_SPHINX_SERVICE_HOST"), question)
}

