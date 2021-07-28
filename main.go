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
		network.ErrorStatus(w)
		return
	}
	defer client.Close()

	token, err := network.GetData("Token", w, r)
	if err != nil {
		return
	}

	user, err := manager.Login(token, r.RemoteAddr, client)
	if err != nil {
		network.ErrorStatus(w)
		return
	}
	tokenJson, err := json.Marshal(user)
	if err != nil {
		network.ErrorStatus(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(tokenJson)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	client, err := dapr.NewClient()
	if err != nil {
		network.ErrorStatus(w)
		return
	}
	defer client.Close()
	token, err := network.GetData("LoginToken", w, r)
	if err != nil {
		return
	}

	user, err := manager.Update(r.RemoteAddr, client, false, token, []byte(""))
	if err != nil {
		network.ErrorStatus(w)
		return
	}
	tokenJson, err := json.Marshal(user)
	if err != nil {
		network.ErrorStatus(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(tokenJson)
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
