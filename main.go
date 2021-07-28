package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/hello-slide/account-manager/manager"
	"github.com/hello-slide/account-manager/network"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
	client, err := dapr.NewClient()
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	defer client.Close()
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	client, err := dapr.NewClient()
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	defer client.Close()

	token := network.Header("Token", r)

	user, err := manager.Login(token[0], r.RemoteAddr, client)
	if err != nil {
		network.ErrorStatus(w)
	}
	tokenJson, err := json.Marshal(user)
	if err != nil {
		network.ErrorStatus(w)
	}
	w.Write([]byte(tokenJson))
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	client, err := dapr.NewClient()
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	defer client.Close()
	token := network.Header("LoginToken", r)

	user, err := manager.Update(r.RemoteAddr, client, false, token[0], []byte(""))
	if err != nil {
		network.ErrorStatus(w)
	}
	tokenJson, err := json.Marshal(user)
	if err != nil {
		network.ErrorStatus(w)
	}
	w.Write([]byte(tokenJson))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/delete", deleteHandler)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
