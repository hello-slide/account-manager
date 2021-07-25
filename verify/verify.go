package verify

import (
	"github.com/futurenda/google-auth-id-token-verifier"
)

func Verify(token string, url string) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	v := googleAuthIDTokenVerifier.Verifier{}
	err := v.VerifyIDToken(token, []string{
		url,
	})
	if err != nil {
		return nil, err
	}
	claimSet, err := googleAuthIDTokenVerifier.Decode(token)
	if err != nil{
		return nil, err
	}
	return claimSet, nil
}
