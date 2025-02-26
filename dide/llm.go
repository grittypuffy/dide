package dide
/*
import (
	"fmt"
	"github.com/rifaideen/talkative"
)


client, err := talkative.New("http://localhost:11434")

if err != nil {
  panic("Failed to create talkative client")
}

model := talkative.DEFAULT_MODEL

// Callback function to handle the response
callback := func(cr *talkative.ChatResponse, err error) {
  if err != nil {
	 fmt.Println(err)
	 return
  }

  fmt.Print(cr.Message.Content)
}

// Additional parameters to include. (Optional)
var params *talkative.ChatParams = nil

// The chat message to send
message := talkative.ChatMessage{
  Role:    talkative.USER, // Initiate the chat as a user
  Content: "What is the capital of France?",
}

// Start the chat
done, err := client.Chat(model, callback, params, message)

if err != nil {
  panic(err)
}

<-done // wait for the chat to complete
fmt.Println()*/