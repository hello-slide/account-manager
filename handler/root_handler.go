package handler

import "net/http"

// Root handler
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi!"))
}
