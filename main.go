package main

import (
	"fmt"
	"github.com/futurenda/google-auth-id-token-verifier"
)

var TOKEN string = ""

func main() {
	v := googleAuthIDTokenVerifier.Verifier{}
	aud := ""
	err := v.VerifyIDToken(TOKEN, []string{
		aud,
	})
	if err == nil {
		claimSet, err := googleAuthIDTokenVerifier.Decode(TOKEN)
		if err != nil {
			fmt.Println("Err")
		}
		fmt.Println(claimSet)
		// claimSet.Iss,claimSet.Email ... (See claimset.go)
	}else{
		fmt.Println("ERR")
	}
}
