package manager

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
)

func Update(ip string, client *dapr.Client, ctx *context.Context, isNew bool, oldToken string, value []byte, seed string) (*ReturnData, error) {
	userTokenState := NewState(client, ctx, LOGIN_TOKEN_STATE)
	newLoginToken, err := CreateLoginToken(ip, client, ctx, seed)
	if err != nil {
		return nil, err
	}
	if isNew {
		// Create a new value
		if err := userTokenState.Set(newLoginToken, value); err != nil {
			return nil, err
		}

	} else {
		// Update
		updateValue, err := userTokenState.Update(oldToken, newLoginToken)
		if err != nil {
			return nil, err
		}
		value = []byte(updateValue)
	}

	sessionToken, err := CreateSessionToken(value, client, ctx)
	if err != nil {
		return nil, err
	}

	return &ReturnData{
		LoginSession: newLoginToken,
		Session:      sessionToken,
	}, nil
}
