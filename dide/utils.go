package dide

import (
	"github.com/rifaideen/talkative"
)

func GetTalkativeClient() *talkative.Client {
	client, err := talkative.New("http://localhost:11434")
	if err != nil {
		panic("No ollama server started")
	}
	return client
}
