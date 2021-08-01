package manager

import (
	"context"

	"github.com/dapr/go-sdk/client"
)

func Logout(ctx *context.Context, client *client.Client, token string) error {
	loginState := NewState(client, ctx, LOGIN_TOKEN_STATE)
	if err := loginState.Delete(token); err != nil {
		return err
	}
	return nil
}
