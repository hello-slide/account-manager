package main

import (
	"fmt"
	"net/http"

	"github.com/hello-slide/account-manager/handler"
	networkutil "github.com/hello-slide/network-util"
)

func init() {
	// initialize dapr client.
	if err := handler.InitClient(); err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.RootHandler)
	mux.HandleFunc("/account/login", handler.LoginHandler)
	mux.HandleFunc("/account/update", handler.UpdateHandler)
	mux.HandleFunc("/account/logout", handler.LogoutHandler)
	mux.HandleFunc("/account/delete", handler.DeleteHandler)

	networkHandler := networkutil.CorsConfig.Handler(mux)

	if err := http.ListenAndServe(":3000", networkHandler); err != nil {
		fmt.Println(err)
	}

	handler.CloseClient()
}
