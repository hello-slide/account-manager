package utils

import (
	"net/http"
	"time"
)

type CookieOperation struct {
	Path     string
	Domain   string
	Secure   bool
	HttpOnly bool
}

func NewCookieOp() *CookieOperation {
	path := "/"
	domain := "hello-slide.jp"

	return &CookieOperation{
		Path:     path,
		Domain:   domain,
		Secure:   true,
		HttpOnly: true,
	}
}

// Set cookie.
//
// Arguments:
//	w {http.ResponseWriter} - http writer.
//	name {string} - cookie key.
//	value {string} - cookie value.
//	exp {int} - date of expiry. It is an hourly unit.
func (c *CookieOperation) Set(w http.ResponseWriter, name string, value string, exp int) {
	expires := time.Now().Add(time.Duration(exp) * time.Hour)
	maxAge := 60 * 60 * exp

	cookie := &http.Cookie{
		Name:  name,
		Value: value,

		Expires: expires,
		MaxAge:  maxAge,

		Secure:   c.Secure,
		Path:     c.Path,
		Domain:   c.Domain,
		HttpOnly: c.HttpOnly,
	}

	http.SetCookie(w, cookie)
}

// Delete cookie
//
// Arguments:
//	w {http.ResponseWriter} - http writer.
//	req {http.Request} - http request.
// name {string} - cookie key.
func (c *CookieOperation) Delete(w http.ResponseWriter, req http.Request, name string) error {
	cookie, err := req.Cookie(name)
	if err != nil {
		return err
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	return nil
}

// Get cookie.
//
// Arguments:
//	req {http.Request} - http request.
//	name {string} - cookie key.
//
// Retruns:
//	{string} - cookie value.
func (c *CookieOperation) Get(req http.Request, name string) (string, error) {
	cookie, err := req.Cookie(name)

	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
