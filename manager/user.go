package manager

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/dapr/go-sdk/client"
	"github.com/hello-slide/account-manager/state"
	oauthapi "google.golang.org/api/oauth2/v2"
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
	if len(value.Value) == 0 {
		return nil, errors.New("user dose not exists")
	}

	var userDataValue *oauthapi.Userinfo = &oauthapi.Userinfo{}

	if err := json.Unmarshal(value.Value, userDataValue); err != nil {
		return nil, err
	}

	selectedData := map[string]string{
		"email":   userDataValue.Email,
		"name":    userDataValue.Name,
		"picture": userDataValue.Picture,
	}

	return json.Marshal(selectedData)
}
