package manager

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

func CreateLoginToken(ip string, client *dapr.Client, ctx *context.Context, seed string) (string, error) {
	var strBuild strings.Builder

	strBuild.WriteString(seed)
	strBuild.WriteString(ip)
	strBuild.WriteString(time.Now().String())

	return createHash([]byte(strBuild.String())), nil
}

func createHash(seed []byte) string {
	result := sha256.Sum256(seed)
	return hex.EncodeToString(result[:])
}
