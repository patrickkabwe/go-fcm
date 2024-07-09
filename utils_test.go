package fcm

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestGenerateGoogleJWT(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	err = privateKey.Validate()
	if err != nil {
		t.Error("Expected no error")
	}
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)

	serviceAccount := &Credentials{
		ClientEmail:  "test",
		TokenURI:     "test",
		PrivateKeyID: "test",
		PrivateKey:   string(keyPEM),
		Type:         "test",
		ProjectID:    "test",
		ClientID:     "test",
		AuthURI:      "test",
	}

	token, err := generateGoogleJWT(serviceAccount)

	if err != nil {
		t.Error("Expected no error")
	}

	if token == "" {
		t.Error("Expected token to be generated")
	}
}
