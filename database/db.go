package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the database connection and creates tables
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:brain.sqlite?_fk=1")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}
	return db, nil
}

func createTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			password TEXT
		);
		CREATE TABLE IF NOT EXISTS things (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT
		);
	`)
	return err
}
