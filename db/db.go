package db

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	pool *pgxpool.Pool

	ErrNoRowsAffected = errors.New("no rows affected")
)

// Connect to database
func Connect() {
	var (
		err error
	)

	dbConnection, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		log.Fatal("db:Connect DATABASE_URL not found")
	}

	// fmt.Sprintf("postgres://%v:%v@%v/%v", db.User, db.Pass, db.Addr, db.DBName))
	pool, err = pgxpool.Connect(context.Background(), dbConnection)
	if err != nil {
		log.Fatal("Connecting database failed:", err)
	}
	log.Println("------------DATABASE IS CONNECTED------------")
}

// Close before exit
func Close() {
	log.Println("db:Close closing db connection")
	pool.Close()
}
