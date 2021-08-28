package handler

import (
	"net/http"

	"github.com/hello-slide/account-manager/manager"
	networkutil "github.com/hello-slide/network-util"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tokenOp := networkutil.NewTokenOp()

	refreshToken, err := tokenOp.GetRefreshToken(r)
	if err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	if err := manager.Logout(ctx, &client, refreshToken); err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	if err := tokenOp.DeleteToken(w, r); err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}
}
