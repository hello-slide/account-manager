package utils

import (
	"context"
	"net/http"
	"strings"

	"github.com/dapr/go-sdk/client"
	networkutil "github.com/hello-slide/network-util"
)

// Get session token.
// If the session token is invalid (expired), redirect to /account/update to update the session token.
//
// Arguments:
//	ctx {context.Context} - context.
//	client {client.Client} - Dapr client.
//	w {http.ResponseWriter} - http writer.
//	r {*http.Request} - http requests.
//	tokenManageName {string} - Dapr app name of token verify.
//	apiUrl {string} - Domain of this API.
//	handlePath {string} - Path of this API.
//
// Returns:
//	{string} - User id.
func GetSessonToken(
	ctx context.Context,
	client client.Client,
	w http.ResponseWriter,
	r *http.Request,
	tokenManageName string,
	apiUrl string,
	handlePath string) (string, error) {

	tokenOp, err := networkutil.NewTokenOp(apiUrl)
	if err != nil {
		return "", err
	}
	sessionToken, err := tokenOp.GetSessionToken(r)
	if err != nil {
		return "", err
	}
	userData, err := VerifySessionToken(ctx, client, sessionToken, tokenManageName)
	if err == nil {
		return userData, nil
	}

	redirectUrl := strings.Join([]string{apiUrl, "/account/update?redirect=", handlePath}, "")

	http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)

	return "", nil
}
