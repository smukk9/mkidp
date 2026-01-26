package main

import (
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/clients", GetClients)

	mux.HandleFunc("POST /oauth/token", tokenEndpoint)
	log.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", mux)
}
