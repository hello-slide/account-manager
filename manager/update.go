package manager

import (
	dapr "github.com/dapr/go-sdk/client"
)

func Update(ip string, client dapr.Client, isNew bool, oldToken string, value []byte) (*ReturnData, error) {
	userTokenState := NewState(client, LOGIN_TOKEN_STATE)
	newLoginToken, err := CreateLoginToken(ip, client)
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
		if err := userTokenState.Update(oldToken, newLoginToken); err != nil {
			return nil, err
		}
	}

	return &ReturnData{
		LoginSession: newLoginToken,
		Session:      "",
	}, nil
}
