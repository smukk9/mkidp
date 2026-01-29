package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	CreatedAt    string `json:"created_at"`
	LastModified string `json:"last_modified"`
}

type UserRequest struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    *string `json:"email"`
}

// UserStore is a store for managing users
type UserStore []User

var UserDB = UserStore{
	{
		ID:           "1",
		Username:     "alice",
		Password:     "password_hash_1",
		Email:        "alice@example.com",
		CreatedAt:    "2025-01-15T08:00:00Z",
		LastModified: "2025-01-15T08:00:00Z",
	},
	{
		ID:           "2",
		Username:     "bob",
		Password:     "password_hash_2",
		Email:        "bob@example.com",
		CreatedAt:    "2025-01-16T12:30:00Z",
		LastModified: "2025-01-16T12:30:00Z",
	},
	{
		ID:           "3",
		Username:     "charlie",
		Password:     "password_hash_3",
		Email:        "charlie@example.com",
		CreatedAt:    "2025-01-17T16:20:00Z",
		LastModified: "2025-01-17T16:20:00Z",
	},
}

//get users

func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UserDB)
}

// add user
func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "Unable to decode the user from the body", http.StatusBadRequest)
	}

	if len(u.Password) < 6 {
		http.Error(w, "Password length should be more that 5", http.StatusBadRequest)
	}

	//TODO: check for user exists

	u.CreatedAt = time.Now().Format(time.RFC3339)
	u.LastModified = time.Now().Format(time.RFC3339)
	u.ID = uuid.New().String()
	UserDB = append(UserDB, User{})
	//for now lets keep this
	json.NewEncoder(w).Encode(u)
}

//get user

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	for _, u := range UserDB {

		if u.ID == id {
			json.NewEncoder(w).Encode(u)
			return
		}
	}

	http.Error(w, "User Not found", http.StatusNotFound)

}
