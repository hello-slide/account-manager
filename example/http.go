package main

import (
	"fmt"
	"net/http"

	"github.com/hello-slide/account-manager/network"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// token := network.Header("Token", r)
	network.GetHeader(w, r)
}

func main() {
	http.HandleFunc("/", rootHandler)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
