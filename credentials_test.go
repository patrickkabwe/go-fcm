package fcm

import "testing"

func TestCredentials_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		credentials *Credentials
	}{
		{
			name: "project_id is required", credentials: &Credentials{},
		},
		{
			name: "private_key is required",
			credentials: &Credentials{
				ProjectID: "project_id",
			},
		},
		{
			name: "client_email is required",
			credentials: &Credentials{
				ProjectID:  "project_id",
				PrivateKey: "private",
			},
		},
		{
			name: "client_id is required",
			credentials: &Credentials{
				ProjectID:   "project_id",
				PrivateKey:  "private",
				ClientEmail: "email",
			},
		},
		{
			name: "auth_uri is required",
			credentials: &Credentials{
				ProjectID:   "project_id",
				PrivateKey:  "private",
				ClientEmail: "email",
				ClientID:    "client_id",
			},
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
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.credentials.Validate()
			if err == nil {
				t.Errorf("expected error, got nil")
			}
		})
	}
}
