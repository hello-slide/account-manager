package manager

import (
	"context"

	"github.com/dapr/go-sdk/client"
	"github.com/hello-slide/account-manager/state"
)

func Delete(ctx *context.Context, client *client.Client, token string) error {
	loginState := state.NewState(client, ctx, refreshTokenState)
	if err := loginState.Delete(token); err != nil {
		return err
	}
	return nil
}
