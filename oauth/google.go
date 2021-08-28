/*
Referred to https://github.com/dapr/components-contrib/blob/master/middleware/http/oauth2/oauth2_middleware.go
*/

package oauth

import (
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

// Get oauth url.
// When authorizing with OAuth, redirect to this URL.
//
// Returns:
//	{string} - Google OAuth url.
func GetAuthURL() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return config.AuthCodeURL(id.String(), oauth2.AccessTypeOffline), nil
}
