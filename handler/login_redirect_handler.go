package handler

import (
	"net/http"
	"strings"

	"github.com/hello-slide/account-manager/manager"
	"github.com/hello-slide/account-manager/oauth"
	"github.com/hello-slide/account-manager/utils"
	networkutil "github.com/hello-slide/network-util"
)

func LoginRedirectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userInfo, err := oauth.Redirect(ctx, r)
	if err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	user, err := manager.Login(ctx, userInfo, r.RemoteAddr, &client)
	if err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	tokenOp, err := networkutil.NewTokenOp(url)
	if err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}
	tokenOp.SetRefreshToken(w, user.RefreshToken)
	tokenOp.SetSessionToken(w, user.Session)

	domain, err := utils.GetDomain(url)
	if err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	http.Redirect(w, r, strings.Join([]string{"https://", domain}, "/dashboard"), http.StatusMovedPermanently)
}
