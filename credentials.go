package fcm

import "fmt"

// Credentials represents the service account credentials required to authenticate with the FCM server.
type Credentials struct {
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

// Validate checks if the required fields are set in the credentials.
func (c *Credentials) Validate() error {
	if c.ProjectID == "" {
		return fmt.Errorf("project_id is required")
	}
	if c.PrivateKey == "" {
		return fmt.Errorf("private_key is required")
	}
	if c.ClientEmail == "" {
		return fmt.Errorf("client_email is required")
	}
	if c.ClientID == "" {
		return fmt.Errorf("client_id is required")
	}
	if c.AuthURI == "" {
		return fmt.Errorf("auth_uri is required")
	}
	if c.TokenURI == "" {
		return fmt.Errorf("token_uri is required")
	}
	if c.AuthProviderX509CertURL == "" {
		return fmt.Errorf("auth_provider_x509_cert_url is required")
	}
	if c.ClientX509CertURL == "" {
		return fmt.Errorf("client_x509_cert_url is required")
	}
	return nil
}
