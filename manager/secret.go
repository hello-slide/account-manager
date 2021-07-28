package manager

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetGoogleOauthPublic() error {
	// opt := map[string]string{
	// 	"version": "2",
	// }
	// secret, err := (*client).GetSecret(*ctx, SECRET_STORE, GOOGLE_OAUTH_PUBLIC_SECRET, opt)
	// if err != nil {
	// 	return err
	// }
	// GoogleOauthKey = secret[GOOGLE_OAUTH_PUBLIC_SECRET]
	// return nil
	url := fmt.Sprintf("http://localhost:3500/v1.0/secrets/%s/%s", SECRET_STORE, GOOGLE_OAUTH_PUBLIC_SECRET)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(res.Body)
	GoogleOauthKey = string(body)
	res.Body.Close()

	return nil
}

func GetSeedValue() error {
	// opt := map[string]string{
	// 	"version": "2",
	// }
	// secret, err := (*client).GetSecret(*ctx, SECRET_STORE, SEED_SECRET, opt)
	// if err != nil {
	// 	return nil
	// }

	// SeedValue = secret[SEED_SECRET]

	// return nil
	url := fmt.Sprintf("http://localhost:3500/v1.0/secrets/%s/%s", SECRET_STORE, SEED_SECRET)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(res.Body)
	SeedValue = string(body)
	res.Body.Close()

	return nil
}
