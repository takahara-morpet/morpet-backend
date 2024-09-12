package config

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	connStr := "user=youruser dbname=takahara-morpet sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Databse connetction error", err)
	}
	log.Println("Databse connected")
}

func CloseDB() {
    if DB != nil {
        err := DB.Close()
        if err != nil {
            log.Println("Failed to close DB", err)
        }
    }
}