package utils

import (
	"context"

	"github.com/dapr/go-sdk/client"
)

func VerifySessionToken(ctx context.Context, daprClient client.Client, token string, tokenManagerName string) (string, error) {
	content := &client.DataContent{
		ContentType: "text/plain",
		Data:        []byte(token),
	}
	responce, err := daprClient.InvokeMethodWithContent(ctx, tokenManagerName, "verify", "post", content)
	if err != nil {
		return "", err
	}
	return string(responce), err
}
