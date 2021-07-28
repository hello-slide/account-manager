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
		userTokenState.Set(newLoginToken, value)

	} else {
		// Update
		userTokenState.Update(oldToken, newLoginToken)
	}

	return &ReturnData{
		LoginSession: newLoginToken,
		Session:      "",
	}, nil
}
