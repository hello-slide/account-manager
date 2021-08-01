package manager

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
)

func Verify(token string, client *dapr.Client, ctx *context.Context, oauthKey string) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	v := googleAuthIDTokenVerifier.Verifier{}
	if err := v.VerifyIDToken(token, []string{oauthKey}); err != nil {
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
