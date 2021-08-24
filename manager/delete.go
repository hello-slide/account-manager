package manager

import (
	"context"
	"encoding/json"

	"github.com/dapr/go-sdk/client"
	"github.com/hello-slide/account-manager/state"
)

func Delete(ctx *context.Context, client *client.Client, token string) error {
	loginState := state.NewState(client, ctx, refreshTokenState)
	userId, err := loginState.Get(token)
	if err != nil {
		return err
	}

	userState := state.NewState(client, ctx, userDataState)

	var userData map[string]string
	_userData, err := userState.Get(string(userId.Value))
	if err != nil {
		return err
	}
	json.Unmarshal(_userData.Value, &userData)

	if err := userState.Delete(string(userId.Value)); err != nil {
		return err
	}

	emailState := state.NewState(client, ctx, userEmailState)
	if err := emailState.Delete(userData["email"]); err != nil {
		return err
	}

	if err := loginState.Delete(token); err != nil {
		return err
	}

	return nil
}
