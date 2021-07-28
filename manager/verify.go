package manager

import (
	"context"
	"fmt"

	dapr "github.com/dapr/go-sdk/client"
	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
)

func Verify(token string, client *dapr.Client, ctx *context.Context) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	publicKey, err := getGoogleOauthPublic(client, ctx)
	if err != nil {
		return nil, err
	}

	v := googleAuthIDTokenVerifier.Verifier{}
	if err := v.VerifyIDToken(token, []string{publicKey}); err != nil {
		return nil, err
	}

	return decode(token)
}

func decode(token string) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	claimSet, err := googleAuthIDTokenVerifier.Decode(token)
	if err != nil {
		return nil, err
	}
	return claimSet, nil
}

func getGoogleOauthPublic(client *dapr.Client, ctx *context.Context) (string, error) {
	opt := map[string]string{
		"version": "2",
	}
	secret, err := (*client).GetSecret(*ctx, SECRET_STORE, GOOGLE_OAUTH_PUBLIC_SECRET, opt)
	if err != nil {
		// return "", err
		return "", fmt.Errorf("getGoogleOauthPublic error")
	}

	return secret[GOOGLE_OAUTH_PUBLIC_SECRET], nil
}
