package token

import "os"

var seed string
var tokenManagerName string
var oauthKey string

func SetEnv() {
	seed = os.Getenv("SEED")
	tokenManagerName = os.Getenv("TOKEN_MANAGER")
	oauthKey = os.Getenv("GOOGLE_OAUTH_KEY")
}
