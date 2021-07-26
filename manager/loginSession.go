package manager

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

func CreateLoginToken(ip string, client dapr.Client) (string, error) {
	var strBuild strings.Builder

	seed, err := getSeed(client)
	if err != nil {
		return "", err
	}
	strBuild.WriteString(seed)
	strBuild.WriteString(ip)
	strBuild.WriteString(time.Now().String())

	return createHash([]byte(strBuild.String())), nil
}

func getSeed(client dapr.Client) (string, error) {
	ctx := context.Background()

	opt := map[string]string{
		"version": "2",
	}
	secret, err := client.GetSecret(ctx, SECRET_STORE, SEED_SECRET, opt)
	if err != nil {
		return "", nil
	}

	return secret[SEED_SECRET], nil
}

func createHash(seed []byte) string {
	result := sha256.Sum256(seed)
	return hex.EncodeToString(result[:])
}
