package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/hello-slide/account-manager/manager"
	"github.com/hello-slide/account-manager/network"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

var client dapr.Client

func loginHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token, err := network.GetData("Token", w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	user, err := manager.Login(token, r.RemoteAddr, &client, &ctx)
	if err != nil {
		network.ErrorStatus(w)
		fmt.Fprintln(w, err)
		return
	}
	tokenJson, err := json.Marshal(user)
	if err != nil {
		network.ErrorStatus(w)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(tokenJson)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token, err := network.GetData("LoginToken", w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	user, err := manager.Update(r.RemoteAddr, &client, &ctx, false, token, []byte(""))
	if err != nil {
		network.ErrorStatus(w)
		fmt.Fprintln(w, err)
		return
	}
	tokenJson, err := json.Marshal(user)
	if err != nil {
		network.ErrorStatus(w)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(tokenJson)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token, err := network.GetData("LoginToken", w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	if err := manager.Logout(&ctx, &client, token); err != nil {
		network.ErrorStatus(w)
		fmt.Fprintln(w, err)
		return
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func init() {
	_client, err := dapr.NewClient()
	if err != nil {
		return
	}
	client = _client

	ctx := context.Background()

	if err := manager.GetGoogleOauthPublic(&client, &ctx); err != nil {
		panic(err)
	}
	if err := manager.GetSeedValue(&client, &ctx); err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/account/login", loginHandler)
	mux.HandleFunc("/account/update", updateHandler)
	mux.HandleFunc("/account/logout", logoutHandler)
	mux.HandleFunc("/account/delete", deleteHandler)

	handler := network.CorsConfig.Handler(mux)

	if err := http.ListenAndServe(":3000", handler); err != nil {
		client.Close()
		fmt.Println(err)
	}
	client.Close()
}
