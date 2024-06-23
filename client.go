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

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type FCMClient struct {
	serviceAccount *ServiceAccount
	httpClient     HttpClient
}

func NewClient() *FCMClient {
	return &FCMClient{httpClient: http.DefaultClient}
}

func (f *FCMClient) Send(msg *MessagePayload) error {

	err := f.makeAPICall(msg)

	if err != nil {
		return err
	}

	return nil
}

func (f *FCMClient) SendToTopic(msg *MessagePayload) error {
	if msg.Message.Topic == "" {
		return fmt.Errorf("topic is required")
	}

	err := f.makeAPICall(msg)

	if err != nil {
		return err
	}

	return nil
}

func (f *FCMClient) SendToCondition() {
	panic("implement me")
}

func (f *FCMClient) SendToMultiple() {
	panic("implement me")
}

func (f *FCMClient) WithCredentialFile(serviceAccountFilePath string) *FCMClient {
	// get current working directory
	dir, err := os.Getwd()

	if err != nil {
		panic("Error getting current working directory")
	}

	serviceAccountFilePath = fmt.Sprintf("%s/%s", dir, serviceAccountFilePath)

	fmt.Println(serviceAccountFilePath)

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

func (f *FCMClient) WithHTTPClient(httpClient HttpClient) *FCMClient {
	f.httpClient = httpClient
	return f
}

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
