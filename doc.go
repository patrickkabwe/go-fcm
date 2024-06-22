// Here is a simple example illustrating how to use FCM library:
//
// func main() {
// client := NewClient().
// 	WithCredentialFile(testServiceAccountFile).
// 	WithHTTPClient(&testHttpClient{
// 		DoFunc: func(req *http.Request) (*http.Response, error) {
// 			return &http.Response{
// 				StatusCode: 200,
// 				Body:       io.NopCloser(bytes.NewReader([]byte(`{}`))),
// 			}, nil
// 		},
// 	})

// err := client.Send(&MessagePayload{
// 	Message: Message{
// 		Token: "test",
// 		Notification: Notification{
// 			Title: "Coming Soon!",
// 			Body:  "Stay tuned for the latest features and improvements. We are constantly working to improve your experience.",
// 		},
// 		Data: map[string]string{
// 			"key": "value",
// 		},
// 		APNS: APNSConfig{
// 			Headers: map[string]string{
// 				"apns-priority": "10",
// 			},
// 			Payload: APNSPayload{
// 				Aps: APNS{
// 					Alert: APNAlert{
// 						Title:    "Coming Soon!",
// 						Subtitle: "This is a subtitle",
// 						Body:     "Stay tuned for the latest features and improvements. We are constantly working to improve your experience.",
// 					},
// 					Badge: 1,
// 					Sound: "default",
// 				},
// 			},
// 		},
// 	},
// })

//	if err != nil {
//		fmt.Println(err)
//		t.Error("Expected no error")
//	}
//
// }
package fcm
