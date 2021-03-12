package models

import "time"

// User - model of user
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
