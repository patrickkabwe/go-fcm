package fcm

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"
)

const (
	testServiceAccountFile = "testdata/service_test.json"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name        string
		serviceFile string
		expectedErr bool
	}{
		{name: "with invalid service file", serviceFile: "invalid.json", expectedErr: true},
		{name: "with valid service file", serviceFile: testServiceAccountFile, expectedErr: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectedErr {
				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered from panic", r)
						if !tc.expectedErr {
							t.Errorf("Expected no panic but got %v", r)
						}
					}
				}()
			}
			client := NewClient().
				WithCredentialFile(tc.serviceFile)

			if !tc.expectedErr && client.serviceAccount == nil {
				t.Error("Expected service account to be set")
			}
		})
	}
}

func TestSend(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr bool
		doFunc      func(req *http.Request) (*http.Response, error)
		payload     *MessagePayload
	}{
		{
			name: "with invalid payload",
			doFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 400,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{"error": {"status": "INVALID_ARGUMENT", "message":"testing"}}`))),
				}, nil
			},
			payload:     &MessagePayload{},
			expectedErr: true,
		},
		{
			name: "with va ",
			doFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{}`))),
				}, nil
			},
			payload: &MessagePayload{
				Message: Message{
					Token: "test",
					Notification: Notification{
						Title: "Coming Soon!",
						Body:  "Stay tuned for the latest features and improvements. We are constantly working to improve your experience.",
					},
					Data: map[string]string{
						"key": "value",
					},
				},
			},
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := NewClient()
			client.WithCredentialFile(testServiceAccountFile).
				WithHTTPClient(&testHttpClient{
					DoFunc: tc.doFunc,
				})
			err := client.Send(tc.payload)

			if tc.expectedErr && err == nil {
				t.Errorf("Expected error but got none")
			} else if !tc.expectedErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
		})
	}
}

func TestSendToTopic(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr bool
		doFunc      func(req *http.Request) (*http.Response, error)
		payload     *MessagePayload
	}{
		{
			name: "with invalid payload",
			doFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 400,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{"error": {"status": "INVALID_ARGUMENT", "message":"testing"}}`))),
				}, nil
			},
			payload:     &MessagePayload{},
			expectedErr: true,
		},
		{
			name: "with valid payload",
			doFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{}`))),
				}, nil
			},
			payload: &MessagePayload{
				Message: Message{
					Topic: "test",
					Notification: Notification{
						Title: "Coming Soon!",
						Body:  "Stay tuned for the latest features and improvements. We are constantly working to improve your experience.",
					},
					Data: map[string]string{
						"key": "value",
					},
				},
			},
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := NewClient()
			client.WithCredentialFile(testServiceAccountFile).
				WithHTTPClient(&testHttpClient{
					DoFunc: tc.doFunc,
				})
			err := client.SendToTopic(tc.payload)

			if tc.expectedErr && err == nil {
				t.Errorf("Expected error but got none")
			} else if !tc.expectedErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
		})
	}
}

func TestSendToCondition(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr bool
		doFunc      func(req *http.Request) (*http.Response, error)
		payload     *MessagePayload
	}{
		{
			name: "with invalid payload",
			doFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 400,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{"error": {"status": "INVALID_ARGUMENT", "message":"testing"}}`))),
				}, nil
			},
			payload:     &MessagePayload{},
			expectedErr: true,
		},
		{
			name: "with valid payload",
			doFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{}`))),
				}, nil
			},
			payload: &MessagePayload{
				Message: Message{
					Condition: "'test' in topics",
					Notification: Notification{
						Title: "Coming Soon!",
						Body:  "Stay tuned for the latest features and improvements. We are constantly working to improve your experience.",
					},
					Data: map[string]string{
						"key": "value",
					},
				},
			},
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := NewClient()
			client.WithCredentialFile(testServiceAccountFile).
				WithHTTPClient(&testHttpClient{
					DoFunc: tc.doFunc,
				})
			err := client.SendToCondition(tc.payload)

			if tc.expectedErr && err == nil {
				t.Errorf("Expected error but got none")
			} else if !tc.expectedErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
		})
	}
}

func TestSendToMultiple(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr bool
		doFunc      func(req *http.Request) (*http.Response, error)
		payload     *MessagePayload
	}{
		{
			name: "with invalid payload",
			doFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 400,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{"error": {"status": "INVALID_ARGUMENT", "message":"testing"}}`))),
				}, nil
			},
			payload:     &MessagePayload{},
			expectedErr: true,
		},
		{
			name: "with valid payload",
			doFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{}`))),
				}, nil
			},
			payload: &MessagePayload{
				Message: Message{
					Tokens: []string{"test1", "test2"},
					Notification: Notification{
						Title: "Coming Soon!",
						Body:  "Stay tuned for the latest features and improvements. We are constantly working to improve your experience.",
					},
					Data: map[string]string{
						"key": "value",
					},
				},
			},
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := NewClient()
			client.WithCredentialFile(testServiceAccountFile).
				WithHTTPClient(&testHttpClient{
					DoFunc: tc.doFunc,
				})
			err := client.SendToMultiple(tc.payload)

			if tc.expectedErr && err == nil {
				t.Errorf("Expected error but got none")
			} else if !tc.expectedErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
		})
	}
}


func TestSendAll(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr bool
		doFunc      func(req *http.Request) (*http.Response, error)
		payload     *MessagePayload
	}{
		{
			name: "with invalid payload",
			doFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 400,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{"error": {"status": "INVALID_ARGUMENT", "message":"testing"}}`))),
				}, nil
			},
			payload:     &MessagePayload{},
			expectedErr: true,
		},
		{
			name: "with token - valid payload",
			doFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{}`))),
				}, nil
			},
			payload: &MessagePayload{
				Message: Message{
					Notification: Notification{
						Title: "Coming Soon!",
						Body:  "Stay tuned for the latest features and improvements. We are constantly working to improve your experience.",
					},
					Data: map[string]string{
						"key": "value",
					},
					Token: "test",
				},
			},
			expectedErr: false,
		},
		{
			name: "with topic - valid payload",
			doFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{}`))),
				}, nil
			},
			payload: &MessagePayload{
				Message: Message{
					Notification: Notification{
						Title: "Coming Soon!",
						Body:  "Stay tuned for the latest features and improvements. We are constantly working to improve your experience.",
					},
					Data: map[string]string{
						"key": "value",
					},
					Topic: "test",
				},
			},
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := NewClient()
			client.WithCredentialFile(testServiceAccountFile).
				WithHTTPClient(&testHttpClient{
					DoFunc: tc.doFunc,
				})
			err := client.SendAll(tc.payload)

			if tc.expectedErr && err == nil {
				t.Errorf("Expected error but got none")
			} else if !tc.expectedErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
		})
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
