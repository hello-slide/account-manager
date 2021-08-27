package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hello-slide/account-manager/manager"
	networkutil "github.com/hello-slide/network-util"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token, err := networkutil.GetFromKey("LoginToken", w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	if err := manager.Logout(&ctx, &client, token); err != nil {
		networkutil.ErrorStatus(w)
		fmt.Fprintln(w, err)
		return
	}
}
