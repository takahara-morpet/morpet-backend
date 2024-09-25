package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        dbURL = "host=db user=me dbname=development sslmode=disable port=5432"
    }

	var err error
	DB, err = sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Database connetction error", err)
	}
	log.Println("Database connected")
}

func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Println("Failed to close DB", err)
		}
	}
}
