package manager

import (
	"context"
	"encoding/json"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/hello-slide/account-manager/state"
	oauthapi "google.golang.org/api/oauth2/v2"
)

func Login(ctx context.Context, userData *oauthapi.Userinfo, ip string, client *dapr.Client) (*ReturnData, error) {
	userId := userData.Id
	userEmail := userData.Email
	userDataJson, err := json.Marshal(userData)
	if err != nil {
		return nil, err
	}

	userState := state.NewState(ctx, client, userDataState)
	if err := userState.Set(userId, []byte(userDataJson)); err != nil {
		return nil, err
	}
	emailState := state.NewState(ctx, client, userEmailState)
	if err := emailState.Set(userEmail, []byte(userId)); err != nil {
		return nil, err
	}

	return Update(ctx, ip, client, true, "", []byte(userId))
}
