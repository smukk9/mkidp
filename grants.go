package main

import (
	"errors"
	"math/rand"
	"time"
)

type Token struct {
	AccessToken string
	IssuedAt    string
	Expiry      string
}

// verifies credentilas and return error or token
func handleClientCredentials(clientId, clientSecret string) (accesstoken string, err error) {

	if !verifyCredentials(clientId, clientSecret) {

		return "", errors.New("Invalid client credentilas")
	}

	return generateToken(clientId), nil

}

func verifyCredentials(clientId, clientSecret string) bool {

	for _, c := range ClientsStore {
		// fmt.Println(c.Id + c.Name)
		if c.Name == clientId && c.Secret == clientSecret {
			return true
		}
	}
	return false

}

func generateToken(clientId string) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 32)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
