package network

import "net/http"

func ErrorStatus(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500"))
}
