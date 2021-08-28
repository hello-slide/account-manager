package oauth

import "os"

var googleClientId string = os.Getenv("GOOGLE_CLIENT_ID")
var googleClientSecret string = os.Getenv("GOOGLE_CLIENT_SECRET")
