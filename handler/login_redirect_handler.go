package handler

import (
	"net/http"

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

	cookieOp := utils.NewCookieOp()
	cookieOp.Set(w, "refresh_token", user.RefreshToken, 24*30)
	cookieOp.Set(w, "session_token", user.Session, 6)
}
