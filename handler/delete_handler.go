package handler

import (
	"fmt"
	"net/http"

	"github.com/hello-slide/account-manager/manager"
	networkutil "github.com/hello-slide/network-util"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tokenOp := networkutil.NewTokenOp()

	refreshToken, err := tokenOp.GetRefreshToken(r)
	if err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	if err := manager.Delete(ctx, &client, refreshToken); err != nil {
		networkutil.ErrorStatus(w)
		fmt.Fprintln(w, err)
		return
	}

	if err := tokenOp.DeleteToken(w, r); err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}
}
