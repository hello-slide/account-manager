package token

import (
	"context"
	"fmt"

	dapr "github.com/dapr/go-sdk/client"
	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
)

func Verify(token string, client *dapr.Client, ctx *context.Context, key string) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	v := googleAuthIDTokenVerifier.Verifier{}
	if err := v.VerifyIDToken(token, []string{key}); err != nil {
		return nil, fmt.Errorf("%v: key: %v", err, key)
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
