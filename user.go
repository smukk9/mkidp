package main

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	CreatedAt    string `json:"created_at"`
	LastModified string `json:"last_modified"`
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

//get user

//add user

//delete user
