package fcm

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateGoogleJWT(serviceAccount *ServiceAccount) (string, error) {
	issuedAt := time.Now().Unix()
	expires := issuedAt + 3600 // 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":   serviceAccount.ClientEmail,
		"sub":   serviceAccount.ClientEmail,
		"aud":   serviceAccount.TokenURI,
		"iat":   issuedAt,
		"exp":   expires,
		"scope": "https://www.googleapis.com/auth/cloud-platform",
	})
	token.Header["kid"] = serviceAccount.PrivateKeyID
	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(serviceAccount.PrivateKey))
	if err != nil {
		return "", err
	}

	return token.SignedString(rsaPrivateKey)
}
