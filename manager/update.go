package manager

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/hello-slide/account-manager/state"
	"github.com/hello-slide/account-manager/utils"
)

func Update(ctx context.Context, ip string, client *dapr.Client, isNew bool, oldToken string, value []byte) (*ReturnData, error) {
	userTokenState := state.NewState(ctx, client, refreshTokenState)
	newRefreshToken, err := utils.CreateRefreshToken(ctx, ip, client, seed)
	if err != nil {
		return nil, err
	}
	if isNew {
		// Create a new value
		// The data will disappear in 30 days(2592000s).
		if err := userTokenState.SetTTL(newRefreshToken, value, "2592000"); err != nil {
			return nil, err
		}

	} else {
		// Update
		// The data will disappear in 30 days(2592000s).
		updateValue, err := userTokenState.Update(oldToken, newRefreshToken, "2592000")
		if err != nil {
			return nil, err
		}
		value = []byte(updateValue)
	}

	sessionToken, err := utils.CreateSessionToken(ctx, value, client, tokenManagerName)
	if err != nil {
		return nil, err
	}

	return &ReturnData{
		RefreshToken: newRefreshToken,
		Session:      sessionToken,
	}, nil
}
