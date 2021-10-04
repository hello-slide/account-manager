package handler

import (
	"net/http"

	"github.com/hello-slide/account-manager/manager"
	"github.com/hello-slide/account-manager/utils"
	networkutil "github.com/hello-slide/network-util"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId, err := utils.GetSessonToken(ctx, client, w, r, tokenManagerName, url, "/account/user")
	if err != nil {
		return
	}

	userData, err := manager.GetUserData(ctx, client, userId)
	if err != nil {
		networkutil.ErrorResponse(w, 2, err)

		// delete cookies
		tokenOp, err := networkutil.NewTokenOp(url)
		if err != nil {
			networkutil.ErrorResponse(w, 1, err)
			return
		}

		if err := tokenOp.DeleteToken(w, r); err != nil {
			networkutil.ErrorResponse(w, 1, err)
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userData)
}
