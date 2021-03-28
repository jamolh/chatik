package db

import (
	"context"
	"log"

	"github.com/jamolh/chatik/models"
)

func GetUser(ctx context.Context, username string) (user models.User, err error) {
	err = pool.QueryRow(ctx, `SELECT 
			id, 
			name,
			last_name, 
			username, 
			password_hash, 
			created_at 
		FROM users 
			WHERE username = $1`, username).Scan(
		&user.ID,
		&user.Name,
		&user.LastName,
		&user.Username,
		&user.Password,
		&user.CreatedAt)
	if err != nil {
		log.Printf("db:GetUser username: %v error: %v\n", username, err)
	}
	return
}

func CheckUserExists(ctx context.Context, username string) (exists bool) {
	err := pool.QueryRow(ctx, `SELECT true FROM users WHERE username = $1`, username).Scan(&exists)
	if err != nil {
		log.Printf("db:UserExists username: %v error: %v", username, err)
	}
	return
}
