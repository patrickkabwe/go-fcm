package fcm

import (
	"testing"
)

func TestCredentials_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		credentials *Credentials
		expectsErr  bool
	}{
		{
			name:        "project_id is required",
			credentials: &Credentials{},
			expectsErr:  true,
		},
		{
			name:        "private_key is required",
			credentials: &Credentials{ProjectID: "project_id"},
			expectsErr:  true,
		},
		{
			name:        "client_email is required",
			credentials: &Credentials{ProjectID: "project_id", PrivateKey: "private"},
			expectsErr:  true,
		},
		{
			name:        "client_id is required",
			credentials: &Credentials{ProjectID: "project_id", PrivateKey: "private", ClientEmail: "email"},
			expectsErr:  true,
		},
		{
			name: "auth_uri is required",
			credentials: &Credentials{
				ProjectID:   "project_id",
				PrivateKey:  "private",
				ClientEmail: "email",
				ClientID:    "client_id",
			},
			expectsErr: true,
		},
		{
			name: "token_uri is required",
			credentials: &Credentials{
				ProjectID:   "project_id",
				PrivateKey:  "private",
				ClientEmail: "email",
				ClientID:    "client_id",
				AuthURI:     "auth_uri",
			},
			expectsErr: true,
		},
		{
			name: "auth_provider_x509_cert_url is required",
			credentials: &Credentials{
				ProjectID:   "project_id",
				PrivateKey:  "private",
				ClientEmail: "email",
				ClientID:    "client_id",
				AuthURI:     "auth_uri",
				TokenURI:    "token_uri",
			},
			expectsErr: true,
		},
		{
			name: "client_x509_cert_url is required",
			credentials: &Credentials{
				ProjectID:               "project_id",
				PrivateKey:              "private",
				ClientEmail:             "email",
				ClientID:                "client_id",
				AuthURI:                 "auth_uri",
				TokenURI:                "token_uri",
				AuthProviderX509CertURL: "auth_provider_x509_cert_url",
			},
			expectsErr: true,
		},
		{
			name: "valid credentials",
			credentials: &Credentials{
				ProjectID:               "project_id",
				PrivateKey:              "private",
				ClientEmail:             "email",
				ClientID:                "client_id",
				AuthURI:                 "auth_uri",
				TokenURI:                "token_uri",
				AuthProviderX509CertURL: "auth_provider_x509_cert_url",
				ClientX509CertURL:       "client_x509_cert_url",
			},
			expectsErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.credentials.Validate()
			if tc.expectsErr && err == nil {
				t.Errorf("expected error, got nil")
			} else if !tc.expectsErr && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}
