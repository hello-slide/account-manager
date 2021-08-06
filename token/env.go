package token

import "os"

var seed string = os.Getenv("SEED")
var tokenManagerName string = os.Getenv("TOKEN_MANAGER")
var oauthKey = os.Getenv("GOOGLE_OAUTH_KEY")
