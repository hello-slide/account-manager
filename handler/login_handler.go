package handler

import (
	"net/http"
	"strings"

	"github.com/hello-slide/account-manager/oauth"
	"github.com/hello-slide/account-manager/utils"
	networkutil "github.com/hello-slide/network-util"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tokenOp, err := networkutil.NewTokenOp(url)
	if err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	// Check the cookie before logging in and try to log in with the refresh token if it exists.
	refreshToken, err := tokenOp.GetRefreshToken(r)
	if err != nil || len(refreshToken) == 0 {
		if err := redirectOAuth(w, r); err != nil {
			networkutil.ErrorResponse(w, 1, err)
			return
		}
	} else {
		_, err = utils.GetSessonToken(ctx, client, w, r, tokenManagerName, url, "/account/login")

		// If for some reason the token verification fails, you can log in again.
		if err != nil {
			if err := redirectOAuth(w, r); err != nil {
				networkutil.ErrorResponse(w, 1, err)
				return
			}
		}

		domain, err := utils.GetDomain(url)
		if err != nil {
			networkutil.ErrorResponse(w, 1, err)
			return
		}
		http.Redirect(w, r, strings.Join([]string{"https://", domain, "/"}, ""), http.StatusMovedPermanently)
	}
}

func redirectOAuth(w http.ResponseWriter, r *http.Request) error {
	url, err := oauth.GetAuthURL()
	if err != nil {
		return err
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
	return nil
}
