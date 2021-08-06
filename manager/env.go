package manager

import "os"

var refreshTokenState = os.Getenv("REFRESH_TOKEN_STATE")
var userDataState = os.Getenv("USER_DATA_STATE")
var userEmailState = os.Getenv("USER_EMAIL_STATE")
