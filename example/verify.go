package main

import (
	"encoding/json"
	"fmt"

	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
)

func Login(token string) {
	publicKey := ""

	claim, err := Verify(token, publicKey)

	// get
	userId := claim.Sub
	fmt.Println(userId)
	if err == nil {
		bytes, _ := json.Marshal(&claim)
		fmt.Println(string(bytes))
	}
}

func Verify(token string, url string) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	v := googleAuthIDTokenVerifier.Verifier{}
	err := v.VerifyIDToken(token, []string{
		url,
	})
	if err != nil {
		return nil, err
	}

	return decode(token)
}

func decode(token string) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	claimSet, err := googleAuthIDTokenVerifier.Decode(token)
	if err != nil {
		return nil, err
	}
	return claimSet, nil
}

func main() {
	Login("")
}
