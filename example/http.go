package main

import (
	"fmt"
	"net/http"

	networkutil "github.com/hello-slide/network-util"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// token := network.Header("Token", r)
	networkutil.GetHeader(w, r)
}

func main() {
	http.HandleFunc("/", rootHandler)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
