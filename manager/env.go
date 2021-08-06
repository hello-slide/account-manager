package manager

import (
	"os"

	"github.com/hello-slide/account-manager/token"
)

var refreshTokenState string
var userDataState string
var userEmailState string

func SetEnv() {
	refreshTokenState = os.Getenv("REFRESH_TOKEN_STATE")
	userDataState = os.Getenv("USER_DATA_STATE")
	userEmailState = os.Getenv("USER_EMAIL_STATE")

	token.SetEnv()
}
