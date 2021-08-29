package oauth

import (
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauthapi "google.golang.org/api/oauth2/v2"
)

var config *oauth2.Config

// Create Google oauth instance.
func SetConfig() {
	url := os.Getenv("API_URL")
	redirect := strings.Join([]string{url, "/account/login/redirect"}, "")

	config = &oauth2.Config{
		ClientID:     googleClientId,
		ClientSecret: googleClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{oauthapi.UserinfoEmailScope, oauthapi.UserinfoProfileScope, oauthapi.OpenIDScope},
		RedirectURL:  redirect,
	}
}
