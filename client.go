package main

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	Id     string
	Name   string
	Secret string
}

type User struct {
	Id       string
	Username string
	Password string
}

// dummy
type ClientStore []Client

type UserStore []User

// Seed initial data
var ClientsStore = ClientStore{
	{Id: "1", Name: "Client_A", Secret: "secret1"},
	{Id: "2", Name: "Client_B", Secret: "secret2"},
	{Id: "3", Name: "Client_C", Secret: "secret3"},
}

var UserDB = UserStore{
	{Id: "1", Username: "username1", Password: "secret1"},
	{Id: "2", Username: "username2", Password: "secret2"},
	{Id: "3", Username: "username3", Password: "secret3"},
}

func GetClients(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ClientsStore)

}
