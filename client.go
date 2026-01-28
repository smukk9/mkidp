package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// TODO: add client grant_type stirng[]
type Client struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Secret       string `json:"secret"`
	RedirectURI  string `json:"redirect_uri"`
	TokenType    string `json:"token_type"`
	TokenExpiry  int    `json:"token_expiry"` // in seconds
	CreatedAt    string `json:"created_at"`
	LastModified string `json:"last_modified"`
	
}

type ClientRequest struct {
	Name        string `json:"name"`
	Secret      string `json:"secret"`
	RedirectURI string `json:"redirect_uri"`
	TokenType   string `json:"token_type"`
	TokenExpiry int    `json:"token_expiry"` // in seconds
	GrantType   string `json:"grant_type"`
}

// ClientStore is a store for managing clients
type ClientStore []Client

// Seed initial data
var clientDB = ClientStore{
	{
		ID:           "1",
		Name:         "Web Application",
		Secret:       "secret_key_1",
		RedirectURI:  "http://localhost:3000/callback",
		TokenType:    "Bearer",
		TokenExpiry:  3600,
		CreatedAt:    "2025-01-20T10:30:00Z",
		LastModified: "2025-01-20T10:30:00Z",
	},
	{
		ID:           "2",
		Name:         "Mobile App",
		Secret:       "secret_key_2",
		RedirectURI:  "mobileapp://callback",
		TokenType:    "Bearer",
		TokenExpiry:  7200,
		CreatedAt:    "2025-01-21T14:15:00Z",
		LastModified: "2025-01-21T14:15:00Z",
	},
	{
		ID:           "3",
		Name:         "Desktop Client",
		Secret:       "secret_key_3",
		RedirectURI:  "http://localhost:5000/auth/callback",
		TokenType:    "Bearer",
		TokenExpiry:  3600,
		CreatedAt:    "2025-01-22T09:45:00Z",
		LastModified: "2025-01-22T09:45:00Z",
	},
}

func GetClients(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientDB)

}

func CreateClient(w http.ResponseWriter, r *http.Request) {

	var req ClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Create new client
	newClient := Client{
		ID:           generateID(), // implement this
		Name:         req.Name,
		Secret:       req.Secret,
		RedirectURI:  req.RedirectURI,
		TokenType:    req.TokenType,
		TokenExpiry:  req.TokenExpiry,
		CreatedAt:    time.Now().Format(time.RFC3339),
		LastModified: time.Now().Format(time.RFC3339),
	}

	clientDB = append(clientDB, newClient)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(newClient)

}

func generateID() string {
	return uuid.New().String()
}
