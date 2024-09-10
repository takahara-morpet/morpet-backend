package config

import (
	"database/sql"
	"log"
	// _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "user=youruser dbname=yourdb sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Databse connetction error", err)
	}
	log.Println("databse connected")
}
