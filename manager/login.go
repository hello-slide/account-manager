package manager

import (
	"encoding/json"

	dapr "github.com/dapr/go-sdk/client"
)

func Login(token string, ip string, client dapr.Client) (*ReturnData, error) {
	claim, err := Verify(token, client)
	if err != nil {
		return nil, err
	}
	userId := claim.Sub
	userDataJson, err := json.Marshal(claim)
	if err != nil {
		return nil, err
	}

	userState := NewState(client, USER_DATA_STATE)
	userState.Set(userId, []byte(userDataJson))

	return Update(ip, client, true, "", []byte(userId))
}
