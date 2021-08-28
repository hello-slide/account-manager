package manager

import (
	"os"
)

var refreshTokenState string = os.Getenv("REFRESH_TOKEN_STATE")
var userDataState string = os.Getenv("USER_DATA_STATE")
var userEmailState string = os.Getenv("USER_EMAIL_STATE")

var seed string = os.Getenv("SEED")
var tokenManagerName string = os.Getenv("TOKEN_MANAGER")
