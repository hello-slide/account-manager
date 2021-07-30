package manager

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
)

func CreateSessionToken(id []byte, client *dapr.Client, ctx *context.Context) (string, error) {
	content := &dapr.DataContent{
		ContentType: "text/plain",
		Data:        id,
	}

	responce, err := (*client).InvokeMethodWithContent(*ctx, TOKEN_MANAGER, "create", "post", content)
	if err != nil {
		return "", nil
	}
	return string(responce), nil
}
