package handler

import (
	"net/http"
	"strings"

	"github.com/hello-slide/account-manager/manager"
	networkutil "github.com/hello-slide/network-util"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	redirectPath := r.URL.Query().Get("redirect")

	ctx := r.Context()

	tokenOp, err := networkutil.NewTokenOp(url)
	if err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	refreshToken, err := tokenOp.GetRefreshToken(r)
	if err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	user, err := manager.Update(ctx, r.RemoteAddr, &client, false, refreshToken, []byte(""))
	if err != nil {
		if _err := tokenOp.DeleteToken(w, r); _err != nil {
			networkutil.ErrorResponse(w, 1, _err)
		} else {
			networkutil.ErrorResponse(w, 1, err)
		}
		return
	}

	tokenOp.SetRefreshToken(w, user.RefreshToken)
	tokenOp.SetSessionToken(w, user.Session)

	w.Header().Set("Cache-Control", "no-cache")

	if len(redirectPath) != 0 {
		redirectUrl := strings.Join([]string{url, redirectPath}, "")
		defer http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
	}
}
