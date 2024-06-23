# FCM Client Library

This library provides a simple interface to interact with Firebase Cloud Messaging (FCM v1) for sending notifications to devices.

## Features

📲 Send messages to devices using FCM.<br>
📢 Send messages to topics.<br>
🔑 Set service account credentials.<br>
🔧 Customize HTTP client for requests.

## Installation

To use this library, simply import it into your Go project:
```go
import "path/to/fcm"
```

Ensure you have the necessary dependencies installed:

```bash
go get -u github.com/patrickkabwe/go-fcm
```

## Usage

### Creating a New Client

To start sending messages, you need to create a new FCM client instance:

```go
client := fcm.NewClient()
```

### Setting Service Account Credentials

Before sending messages, set your service account credentials:

```go
client = client.WithCredentialFile("path/to/serviceAccountKey.json")
```

### Sending a Message

To send a message, create a `MessagePayload` and use the `Send` method:

```go
msg := &MessagePayload{
    // Populate your message payload
}
err := client.Send(msg)
if err != nil {
    log.Fatalf("Failed to send message: %v", err)
}

log.Println("Message sent successfully")
```

### Sending a Message to a Topic

To send a message to a specific topic:

```go
msg := &MessagePayload{
    // Populate your message payload
    Topic: "news",
}

err := client.SendToTopic(msg)
if err != nil {
    log.Fatalf("Failed to send message: %v", err)
}

log.Println("Message sent successfully")
```

### Customizing HTTP Client

You can customize the HTTP client used for making requests:

```go
customClient := &http.Client{Timeout: time.Second * 10}
client = client.WithHTTPClient(customClient)
```

## Contributing
Contributions to this library are welcome. Please ensure to follow the coding standards and write tests for new features.

## License
This library is licensed under the MIT License. See the [LICENSE](/LICENCE) file for details.