package manager

import (
	"context"
	"encoding/json"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/hello-slide/account-manager/state"
	_token "github.com/hello-slide/account-manager/token"
)

func Login(token string, ip string, client *dapr.Client, ctx *context.Context, oauthKey string, seed string) (*ReturnData, error) {
	claim, err := _token.Verify(token, client, ctx)
	if err != nil {
		return nil, err
	}
	userId := claim.Sub
	userEmail := claim.Email
	userDataJson, err := json.Marshal(claim)
	if err != nil {
		return nil, err
	}

	userState := state.NewState(client, ctx, userDataState)
	if err := userState.Set(userId, []byte(userDataJson)); err != nil {
		return nil, err
	}
	emailState := state.NewState(client, ctx, userEmailState)
	if err := emailState.Set(userEmail, []byte(userId)); err != nil {
		return nil, err
	}

	return Update(ip, client, ctx, true, "", []byte(userId), seed)
}
