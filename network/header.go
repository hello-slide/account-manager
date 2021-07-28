package network

import (
	"net/http"
)

func Header(tag string, r *http.Request) []string {
	h := r.Header[tag]
	return h
}
