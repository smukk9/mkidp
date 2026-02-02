package main

import (
	"crypto"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/square/go-jose.v2/jwt"
)

//verify dpop header

func HandelDpopTokenRequest(r *http.Request) (string, error) {

	dpopHearder := r.Header.Get("dpop")

	// fmt.Println("Recevied a dpop header: " + dpopHearder)
	token, err := jwt.ParseSigned(dpopHearder)

	if err != nil {
		return "", errors.New("Invalid DPOP JWT format")
	}

	var claims map[string]interface{}

	//TODO: verify claims match the request
	err = token.UnsafeClaimsWithoutVerification(&claims)
	if err != nil {
		return "", errors.New("Unable to calculate thumbrpint, invaild JsonWebKey")
	}

	// //figure out what token has
	// for _, v := range token.Headers {
	// 	fmt.Print("----------")
	// 	fmt.Print(v)
	// 	fmt.Print("----------")
	// }

	//get the header[0] that has the JsonWebKey
	//must contiain only one. Pick the first one.
	jk := token.Headers[0].JSONWebKey

	jkt, err := jk.Thumbprint(crypto.SHA256)
	if err != nil {
		return "", errors.New("Unable to calculate thumbrpint, invaild JsonWebKey")
	}
	//https://datatracker.ietf.org/doc/html/rfc9449#name-jwk-thumbprint-confirmation
	//JWK SHA-256 Thumbprint confirmation method. The value of the jkt member MUST be the base64url encoding (as defined in [RFC7515]) of the JWK SHA-256 Thumbprint (according to [RFC7638]) of the DPoP public key (in JWK format) to which the access token is bound.
	fmt.Println("----------")
	fmt.Println(base64.RawURLEncoding.EncodeToString(jkt))
	fmt.Println("----------")

	fmt.Println("Subject:", claims["sub"])
	fmt.Println("HTM:", claims["htm"])
	fmt.Println("HTU:", claims["htu"])

	return base64.RawURLEncoding.EncodeToString(jkt), nil

}
