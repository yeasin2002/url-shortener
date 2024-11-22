package utils

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)



func ConnectDB(dbFile string) *sql.DB {
	
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	// Test the database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v\n", err)
	}

	log.Println("Connected to SQLite database!")
	return db
}
