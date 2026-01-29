package main

import (
	"fmt"
	"net/http"

	"github.com/square/go-jose/jwt"
)

//verify dpop header

func HandelDpopTokenRequest(r *http.Request) {

	dpopHearder := r.Header.Get("dpop")

	fmt.Println("Recevied a dpop header: " + dpopHearder)
	token, err := jwt.ParseSigned(dpopHearder)

	if err != nil {
		fmt.Println("Invalid DPoP JWT format:", err)
		return
	}

	fmt.Printf("DPoP Headers: %+v\n", token.Headers)

}
