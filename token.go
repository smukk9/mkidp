package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TokenDB struct {
	AccessToken string
	ExpireAt    string
	CreatedAt   string
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	Expiry      string `json:"expiry"`
	TokenType   string `json:"token_type"`
}

func tokenEndpoint(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	clientId := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")
	grantType := r.FormValue("grant_type")

	fmt.Println(grantType)

	var err error
	var access_token string

	switch grantType {

	case "client_credentials":

		access_token, err = handleClientCredentials(clientId, clientSecret)

	case "password":

		username := r.FormValue("username")
		password := r.FormValue("password")
		access_token, err = handlePasswordGrantType(clientId, clientSecret, username, password)

	default:
		http.Error(w, "Unsupported grant type", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Invalid client credentilas", http.StatusUnauthorized)
		return
	}

	var res TokenResponse
	res.AccessToken = access_token
	res.TokenType = "Bearer"
	res.Expiry = "299"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
