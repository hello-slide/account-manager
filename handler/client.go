package handler

import (
	dapr "github.com/dapr/go-sdk/client"
)

var client dapr.Client

// Initialize dapr client.
func InitClient() error {
	_client, err := dapr.NewClient()
	if err != nil {
		return err
	}
	client = _client

	return nil
}

// Close dapr client
func CloseClient() {
	client.Close()
}
