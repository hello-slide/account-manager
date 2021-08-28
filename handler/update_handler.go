package handler

import (
	"fmt"
	"net/http"

	"github.com/hello-slide/account-manager/manager"
	networkutil "github.com/hello-slide/network-util"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tokenOp := networkutil.NewTokenOp()

	refreshToken, err := tokenOp.GetRefreshToken(r)
	if err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	user, err := manager.Update(ctx, r.RemoteAddr, &client, false, refreshToken, []byte(""))
	if err != nil {
		networkutil.ErrorStatus(w)
		fmt.Fprintln(w, err)
		return
	}

	tokenOp.SetRefreshToken(w, user.RefreshToken)
	tokenOp.SetSessionToken(w, user.Session)
}
