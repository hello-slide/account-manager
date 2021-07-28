package manager

import (
	"context"
	"encoding/json"

	dapr "github.com/dapr/go-sdk/client"
)

func Login(token string, ip string, client *dapr.Client, ctx *context.Context) (*ReturnData, error) {
	claim, err := Verify(token, client, ctx)
	if err != nil {
		return nil, err
	}
	userId := claim.Sub
	userDataJson, err := json.Marshal(claim)
	if err != nil {
		return nil, err
	}

	userState := NewState(client, ctx, USER_DATA_STATE)
	if err := userState.Set(userId, []byte(userDataJson)); err != nil {
		return nil, err
	}

	return Update(ip, client, ctx, true, "", []byte(userId))
}
