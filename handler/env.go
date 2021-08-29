package handler

import (
	"os"
)

var url = os.Getenv("API_URL")
var tokenManagerName string = os.Getenv("TOKEN_MANAGER")
