package manager

import (
	"context"

	"github.com/dapr/go-sdk/client"
	"github.com/hello-slide/account-manager/state"
)

// Get user data.
//
// Arguments:
//	ctx {context.Context} - context
//	client {client.Client} - dapr client.
//	userId {string} - user id.
//
// Returns:
//	{[]byte} - user data.
func GetUserData(ctx context.Context, client client.Client, userId string) ([]byte, error) {
	userData := state.NewState(ctx, &client, userDataState)

	value, err := userData.Get(userId)
	if err != nil {
		return nil, err
	}

	return value.Value, nil
}
