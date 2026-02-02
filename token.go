package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// Initialize logger at package level
var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type TokenRecord struct {
	AccessToken     string
	Thumbprint_DPOP string
	ExpireAt        int64
	CreatedAt       int64
	Client_id       string
	GrantType       string
	TokenType       string
}

var tokenStore []TokenRecord

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	Expiry      int64  `json:"expiry"`
	TokenType   string `json:"token_type"`
}

func tokenEndpoint(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	//TODO: 1. move this to extracDPOPTHumprint emthods
	var jkt string
	if r.Header.Get("dpop") != "" {
		var err error
		jkt, err = HandelDpopTokenRequest(r)
		if err != nil {
			http.Error(w, "dpop error", http.StatusBadRequest)
		}
	}
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

	var ts TokenRecord
	ts.AccessToken = access_token
	ts.CreatedAt = time.Now().Unix()
	ts.ExpireAt = time.Now().Unix() + 300

	if r.Header.Get("dpop") != "" && jkt != "" {
		ts.Thumbprint_DPOP = jkt
		println(jkt)
		ts.TokenType = "DPOP"
	} else {
		ts.Thumbprint_DPOP = ""
		ts.TokenType = "Bearer"
	}

	ts.Client_id = clientId
	ts.GrantType = grantType
	tokenStore = append(tokenStore, ts)
	fmt.Println("--------------------------")
	printTokenStore(tokenStore)
	fmt.Println("--------------------------")

	var res TokenResponse
	res.AccessToken = ts.AccessToken
	res.TokenType = ts.TokenType
	res.Expiry = ts.ExpireAt

	logger.Info("token store:", "count", len(tokenStore), "tokens", tokenStore)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// for debuging only
// TODO:L: remove
func printTokenStore(tokens []TokenRecord) {
	for i, t := range tokens {
		fmt.Printf("Record %d:\n", i+1)
		fmt.Printf("  AccessToken: %s\n", t.AccessToken)
		fmt.Printf("  Thumbprint_DPOP: %s\n", t.Thumbprint_DPOP)
		fmt.Printf("  ExpireAt: %d\n", t.ExpireAt)
		fmt.Printf("  CreatedAt: %d\n", t.CreatedAt)
		fmt.Println()
	}
}

// TODO: clean up from above and put it here
// func generateAccessToken()

// TODO; createTokenRecord
// func createTokenRecord(accessToken, clientID, grantType, jkt string) TokenRecord {}
