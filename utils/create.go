package utils

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
)

func CreateSessionToken(ctx context.Context, id []byte, client *dapr.Client, tokenManagerName string) (string, error) {
	content := &dapr.DataContent{
		ContentType: "text/plain",
		Data:        id,
	}

	responce, err := (*client).InvokeMethodWithContent(ctx, tokenManagerName, "create", "post", content)
	if err != nil {
		return "", err
	}
	return string(responce), nil
}
