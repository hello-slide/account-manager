package utils

import (
	_url "net/url"
	"strings"
)

// Get domain from url.
//
// Example:
//	https://a.example.com/aa/bb -> example.com
//
// Arguments:
//	url {string} - target url.
//
// Returns:
//	{string} - domain value.
func GetDomain(url string) (string, error) {
	u, err := _url.Parse(url)
	if err != nil {
		return "", err
	}
	splittedHost := strings.Split(u.String(), ".")
	hostLen := len(splittedHost)
	domain := strings.Join(splittedHost[hostLen-2:], ".")

	return domain, nil
}
