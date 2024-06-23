package main

import (
	"fmt"
	"log"

	"github.com/patrickkabwe/go-fcm"
)

func main() {
	// Create a new FCM client.
	client := fcm.NewClient().
		WithCredentialFile("path/to/your/credential.json")

	// Create a new message payload.
	msg := &fcm.MessagePayload{
		Message: fcm.Message{
			Notification: fcm.Notification{
				Title: "Hello",
				Body:  "World",
			},
			Topic: "news",
		},
	}

	// Send the message to the FCM server.
	if err := client.SendToTopic(msg); err != nil {
		log.Fatalf("failed to send message: %v", err)
	}

	fmt.Println("message sent successfully")

}
