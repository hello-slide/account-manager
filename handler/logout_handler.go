package handler

import (
	"net/http"

	"github.com/hello-slide/account-manager/manager"
	"github.com/hello-slide/account-manager/utils"
	networkutil "github.com/hello-slide/network-util"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cookieOp := utils.NewCookieOp()

	refreshToken, err := cookieOp.Get(r, "refresh_token")
	if err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	if err := manager.Logout(ctx, &client, refreshToken); err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	if err := cookieOp.Delete(w, r, "session_token"); err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	if err := cookieOp.Delete(w, r, "refresh_token"); err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}
}
