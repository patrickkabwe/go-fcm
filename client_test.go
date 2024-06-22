package fcm

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"
)

const (
	testServiceAccountFile = "testdata/service_account.json"
)

func TestNew(t *testing.T) {
	client := NewClient().
		WithCredentialFile(testServiceAccountFile)
	if client == nil {
		t.Error("Expected client to be created")
	}
}

func TestSend_WithNoMessage(t *testing.T) {
	client := NewClient()
	client.WithCredentialFile(testServiceAccountFile).
		WithHTTPClient(&testHttpClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 400,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{"error": {"status": "INVALID_ARGUMENT", "message":"testing"}}`))),
				}, errors.New("error")
			},
		})
	err := client.Send(&MessagePayload{})

	if err == nil {
		t.Error("Expected error")
	}
}

func TestSend_WithMessage(t *testing.T) {
	client := NewClient().
		WithCredentialFile(testServiceAccountFile).
		WithHTTPClient(&testHttpClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{}`))),
				}, nil
			},
		})

	err := client.Send(&MessagePayload{
		Message: Message{
			Token: "test",
			Notification: Notification{
				Title: "New iLunge Version Coming Soon!",
				Body:  "Stay tuned for the latest features and improvements. We are constantly working to improve your experience.",
			},
			Data: map[string]string{
				"key": "value",
			},
			APNS: APNSConfig{
				Headers: map[string]string{
					"apns-priority": "10",
				},
				Payload: APNSPayload{
					Aps: APNS{
						Alert: APNAlert{
							Title:    "New iLunge Version Coming Soon!",
							Subtitle: "This is a subtitle",
							Body:     "Stay tuned for the latest features and improvements. We are constantly working to improve your experience.",
						},
						Badge: 1,
						Sound: "default",
					},
				},
			},
		},
	})

	if err != nil {
		fmt.Println(err)
		t.Error("Expected no error")
	}
}

func TestGetAccessToken(t *testing.T) {
	resBody := `{"access_token":"test","expires_in":3600}`
	client := NewClient().
		WithCredentialFile(testServiceAccountFile).
		WithHTTPClient(&testHttpClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader([]byte(resBody))),
				}, nil
			},
		})
	token := client.getAccessToken(client.serviceAccount)

	if token == "" {
		t.Error("Expected token to be generated")
	}
}

type testHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (t *testHttpClient) Do(req *http.Request) (*http.Response, error) {
	if t.DoFunc != nil {
		return t.DoFunc(req)
	}
	return &http.Response{}, nil
}
