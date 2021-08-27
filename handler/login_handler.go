package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hello-slide/account-manager/manager"
	networkutil "github.com/hello-slide/network-util"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token, err := networkutil.GetFromKey("Token", w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	user, err := manager.Login(token, r.RemoteAddr, &client, &ctx)
	if err != nil {
		networkutil.ErrorStatus(w)
		fmt.Fprintln(w, err)
		return
	}
	tokenJson, err := json.Marshal(user)
	if err != nil {
		networkutil.ErrorStatus(w)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(tokenJson)
}
