package manager

import (
	"context"
	dapr "github.com/dapr/go-sdk/client"
)

func Login(token string, client dapr.Client) (string, error) {
	publicKey, err := getGoogleOauthPublic(client)
	if err != nil {
		return "", err
	}
	claim, err := Verify(token, publicKey)
}


func getGoogleOauthPublic(client dapr.Client) (string, error) {
	ctx := context.Background()

	opt := map[string]string{
		"version": "2",
	}
	secret, err := client.GetSecret(ctx, SECRET_STORE, GOOGLE_OAUTH_PUBLIC_SECRET, opt)
	if err != nil {
		return "", nil
	}

	return secret[GOOGLE_OAUTH_PUBLIC_SECRET], nil
}
