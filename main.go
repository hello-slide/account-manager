package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/hello-slide/account-manager/manager"
	networkutil "github.com/hello-slide/network-util"
)

var client dapr.Client

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(manager.GetKey()))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
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

func updateHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token, err := networkutil.GetFromKey("LoginToken", w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	user, err := manager.Update(r.RemoteAddr, &client, &ctx, false, token, []byte(""))
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

func logoutHandler(w http.ResponseWriter, r *http.Request) {
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

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token, err := networkutil.GetFromKey("LoginToken", w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	if err := manager.Delete(&ctx, &client, token); err != nil {
		networkutil.ErrorStatus(w)
		fmt.Fprintln(w, err)
		return
	}
}

func init() {
	_client, err := dapr.NewClient()
	if err != nil {
		return
	}
	client = _client
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/account/login", loginHandler)
	mux.HandleFunc("/account/update", updateHandler)
	mux.HandleFunc("/account/logout", logoutHandler)
	mux.HandleFunc("/account/delete", deleteHandler)

	handler := networkutil.CorsConfig.Handler(mux)

	if err := http.ListenAndServe(":3000", handler); err != nil {
		client.Close()
		fmt.Println(err)
	}
	client.Close()
}
