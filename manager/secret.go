package manager

import (
	"context"

	"github.com/dapr/go-sdk/client"
)

func GetGoogleOauthPublic(client *client.Client, ctx *context.Context) error {
	opt := map[string]string{
		"version": "2",
	}
	secret, err := (*client).GetSecret(*ctx, SECRET_STORE, GOOGLE_OAUTH_PUBLIC_SECRET, opt)
	if err != nil {
		return err
	}
	GoogleOauthKey = secret[GOOGLE_OAUTH_PUBLIC_SECRET]
	return nil
}

func GetSeedValue(client *client.Client, ctx *context.Context) error {
	opt := map[string]string{
		"version": "2",
	}
	secret, err := (*client).GetSecret(*ctx, SECRET_STORE, SEED_SECRET, opt)
	if err != nil {
		return nil
	}

	SeedValue = secret[SEED_SECRET]

	return nil
}
