package main

import (
	"errors"
	"fmt"
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

	//TODO ensure the token grant_type as well by getting the clientDB data

	return generateToken(clientId), nil

}

func verifyCredentials(clientId, clientSecret string) bool {

	for _, c := range clientDB {
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

// verifies credentilas and return error or token
func handlePasswordGrantType(clientId, clientSecret, username, password string) (accesstoken string, err error) {

	if !verifyCredentials(clientId, clientSecret) {
		fmt.Println("invalid credentilas")
		return "", errors.New("Invalid client credentilas")
	}

	if !verifyUsernamePassword(username, password) {
		fmt.Println("invalid username")
		return "", errors.New("invalid username and password")
	}

	return generateToken(clientId), nil

}

func verifyUsernamePassword(username, password string) bool {

	for _, c := range UserDB {
		// fmt.Println(c.Id + c.Name)
		if c.Username == username && c.Password == password {
			return true
		}
	}
	return false

}
