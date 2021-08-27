package oauth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauthapi "google.golang.org/api/oauth2/v2"
)

var config *oauth2.Config

// Create Google oauth instance.
//
// Arguments:
//	redirect {string} - redirect url.
func SetConfig(redirect string) {
	config = &oauth2.Config{
		ClientID:     googleClientId,
		ClientSecret: googleClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{oauthapi.UserinfoEmailScope},
		RedirectURL:  redirect,
	}
}
