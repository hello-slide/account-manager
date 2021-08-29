package oauth

import (
	"context"
	"fmt"
	"net/http"

	oauthapi "google.golang.org/api/oauth2/v2"
)

// Google OAuth redirect.
//
// Arguments:
//	ctx {context.Context} - Context.
//	r {*http.Request} - request.
//
// Returns:
//	{*oauthapi.Userinfo} - User info.
func Redirect(ctx context.Context, r *http.Request) (*oauthapi.Userinfo, error) {
	code := r.URL.Query()["code"]
	if len(code) == 0 {
		return nil, fmt.Errorf("invald parameter")
	}

	token, err := config.Exchange(ctx, code[0])
	if err != nil {
		return nil, err
	}

	client := config.Client(ctx, token)
	svr, err := oauthapi.New(client)

	if err != nil {
		return nil, err
	}
	userInfo, err := svr.Userinfo.Get().Do()
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
