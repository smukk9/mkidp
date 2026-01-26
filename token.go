package main

import (
	"encoding/json"
	"net/http"
)

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

	if grantType != "client_credentials" {
		http.Error(w, "Unsupported grant type", http.StatusBadRequest)
		return
	}

	atoken, err := handleClientCredentials(clientId, clientSecret)

	if err != nil {
		http.Error(w, "Invalid clinet credentilas", http.StatusUnauthorized)
		return
	}

	var res TokenResponse
	res.AccessToken = atoken
	res.TokenType = "Bearer"
	res.Expiry = "299"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
