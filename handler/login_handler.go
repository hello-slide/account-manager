package handler

import (
	"net/http"

	"github.com/hello-slide/account-manager/oauth"
	networkutil "github.com/hello-slide/network-util"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	url, err := oauth.GetAuthURL()
	if err != nil {
		networkutil.ErrorResponse(w, 1, err)
		return
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
