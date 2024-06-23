package fcm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	FCM_V1_URL = "https://fcm.googleapis.com/v1/projects/%s/messages:send"
	SCOPES     = "https://www.googleapis.com/auth/firebase.messaging"
)

// ServiceAccount represents the credentials for a service account.
type ServiceAccount struct {
	Type                    string `json:"type,omitempty"`
	ProjectID               string `json:"project_id,omitempty"`
	PrivateKeyID            string `json:"private_key_id,omitempty"`
	PrivateKey              string `json:"private_key,omitempty"`
	ClientEmail             string `json:"client_email,omitempty"`
	ClientID                string `json:"client_id,omitempty"`
	AuthURI                 string `json:"auth_uri,omitempty"`
	TokenURI                string `json:"token_uri,omitempty"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url,omitempty"`
	ClientX509CertURL       string `json:"client_x509_cert_url,omitempty"`
}

// HttpClient is an interface that represents an HTTP client.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// FCMClient represents a client for interacting with the Firebase Cloud Messaging (FCM) service.
type FCMClient struct {
	serviceAccount *ServiceAccount
	httpClient     HttpClient
}

// NewClient creates a new FCMClient instance with the default HTTP client.
// This uses the new version of FCM API (v1) to send messages to devices.
func NewClient() *FCMClient {
	return &FCMClient{httpClient: http.DefaultClient}
}

// Send sends the given message payload to the FCM server.
// It returns an error if the API call fails.
func (f *FCMClient) Send(msg *MessagePayload) error {
	return f.makeAPICall(msg)
}

// SendToTopic sends a message payload to a specific topic.
// It returns an error if the topic is empty or if there was an error making the API call.
func (f *FCMClient) SendToTopic(msg *MessagePayload) error {
	if msg.Message.Topic == "" {
		return fmt.Errorf("topic is required")
	}
	return f.makeAPICall(msg)
}

// SendToCondition sends a message to devices that match the given condition.
func (f *FCMClient) SendToCondition() {
	panic("implement me")
}

// SendToMultiple sends a message to multiple devices using the FCM (Firebase Cloud Messaging) service.
func (f *FCMClient) SendToMultiple() {
	panic("implement me")
}

// WithCredentialFile sets the service account credentials for the FCM client
// by reading the credentials from the specified file path.
// It returns the modified FCMClient instance.
// If the service account file is not found or there is an error parsing the file,
// it will panic with an appropriate error message.
func (f *FCMClient) WithCredentialFile(serviceAccountFilePath string) *FCMClient {
	file, err := os.ReadFile(serviceAccountFilePath)
	if os.IsNotExist(err) {
		panic("Service account file not found")
	}

	var serviceAccount ServiceAccount

	err = json.Unmarshal(file, &serviceAccount)

	if err != nil {
		panic("Error parsing service account file")
	}

	f.serviceAccount = &serviceAccount

	return f
}

// WithHTTPClient sets the HTTP client to be used by the FCM client.
// It allows you to customize the HTTP client used for making requests to the FCM server.
// The provided httpClient should implement the HttpClient interface.
// Returns the FCM client itself to allow for method chaining.
func (f *FCMClient) WithHTTPClient(httpClient HttpClient) *FCMClient {
	f.httpClient = httpClient
	return f
}

// makeAPICall sends an HTTP POST request to the FCM API with the provided message payload.
// It marshals the message payload into JSON format and includes it in the request body.
// The function sets the necessary headers, makes the API request, and handles the response.
// If any error occurs during the process, it is returned.
func (f *FCMClient) makeAPICall(msg *MessagePayload) error {
	jsonData, err := json.Marshal(msg)

	if err != nil {
		log.Println("Error marshalling message payload")
		return err
	}
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(FCM_V1_URL, f.serviceAccount.ProjectID),
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		log.Println("Error creating API request object", err)
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", f.getAccessToken(f.serviceAccount)))
	res, err := f.httpClient.Do(req)

	if err != nil {
		log.Println("Error making API request", err)
		return err
	}

	defer res.Body.Close()

	return f.handleResponse(res)
}

// getAccessToken generates and retrieves an access token for the FCM client using the provided service account.
// It first generates a Google JWT using the given service account, then uses the JWT to obtain an access token
// from Google. If any error occurs during the process, an empty string is returned.
func (f *FCMClient) getAccessToken(serviceAccount *ServiceAccount) string {
	jwt, err := generateGoogleJWT(serviceAccount)

	if err != nil {
		return ""
	}
	token, err := f.getAccessTokenFromGoogle(jwt)
	if err != nil {
		return ""
	}
	return token
}

// getAccessTokenFromGoogle retrieves an access token from Google using the provided JWT.
// It sends a POST request to the TokenURI endpoint of the service account with the JWT as the assertion.
// The function returns the access token as a string if successful, otherwise it returns an error.
func (f *FCMClient) getAccessTokenFromGoogle(jwt string) (string, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		f.serviceAccount.TokenURI,
		bytes.NewBuffer([]byte(fmt.Sprintf("grant_type=urn:ietf:params:oauth:grant-type:jwt-bearer&assertion=%s", jwt))),
	)

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := f.httpClient.Do(req)

	if err != nil {
		return "", err
	}

	var response map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return "", err
	}

	accessToken, ok := response["access_token"].(string)

	if !ok {
		return "", fmt.Errorf("failed to get access token")
	}

	return accessToken, nil
}

// handleResponse decodes the response body from an HTTP response and handles the FCM server's response.
// It returns an error if there was an error decoding the response body or if the FCM server returned an error status.
// If the response is successful, it logs a success message and returns nil.
func (f *FCMClient) handleResponse(res *http.Response) error {
	var response map[string]interface{}

	err := json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		log.Println("Error decoding response body", err)
		return err
	}
	switch res.StatusCode != http.StatusOK {
	case true:
		status := response["error"].(map[string]interface{})["status"].(string)
		message := response["error"].(map[string]interface{})["message"].(string)
		log.Println(status, message)
		return fmt.Errorf(`%s: %s`, status, message)
	default:
		log.Println("Message sent successfully")
		return nil
	}
}
