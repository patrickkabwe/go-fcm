package main

import (
	"fmt"
	"log"

	"github.com/patrickkabwe/go-fcm"
)

func main() {
	// Create a new FCM client.
	client := fcm.NewClient().
		WithCredentialFile("service_test.json")

	// Create a new message payload.
	msg := &fcm.MessagePayload{
		Message: fcm.Message{
			Notification: fcm.Notification{
				Title: "Hello",
				Body:  "World",
			},
			Token: "fRo..............THx",
		},
	}

	// Send the message to the FCM server.
	if err := client.Send(msg); err != nil {
		log.Fatalf("failed to send message: %v", err)
	}

	fmt.Println("message sent successfully")

}
