package manager

import (
	"context"

	"github.com/dapr/go-sdk/client"
	"github.com/hello-slide/account-manager/state"
)

func Logout(ctx context.Context, client *client.Client, token string) error {
	loginState := state.NewState(ctx, client, refreshTokenState)
	if err := loginState.Delete(token); err != nil {
		return err
	}
	return nil
}
