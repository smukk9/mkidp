package main

import (
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/client", GetClients)
	mux.HandleFunc("POST /api/v1/client", CreateClient)

	mux.HandleFunc("POST /oauth/token", tokenEndpoint)
	log.Println("Server starting on :8098...")
	http.ListenAndServe(":8098", mux)
}
