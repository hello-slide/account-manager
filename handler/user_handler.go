package handler

import (
	"net/http"
	"strings"

	"github.com/hello-slide/account-manager/manager"
	"github.com/hello-slide/account-manager/utils"
	networkutil "github.com/hello-slide/network-util"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tokenOp, err := networkutil.NewTokenOp(url)
	if err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	sessionToken, err := tokenOp.GetSessionToken(r)
	if err != nil {
		return
	}

	userId, err := utils.VerifySessionToken(ctx, client, sessionToken, tokenManagerName)

	if err != nil {
		redirectUrl := strings.Join([]string{url, "/account/update?redirect=/account/user"}, "")

		http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
		return
	}

	userData, err := manager.GetUserData(ctx, client, userId)
	if err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userData)
}