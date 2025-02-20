package database

import (
	"context"
	"database/sql"
	"fmt"
	"go-brain/internal/model"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

// InitDB initializes the database connection and creates tables if they don't exist
func InitDB() (*bun.DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:brain.sqlite?_fk=1")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite database: %w", err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())

	// AutoMigrate to create tables
	if err := autoMigrate(db); err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %w", err)
	}

	return db, nil
}

func autoMigrate(db *bun.DB) error {
	ctx := context.Background()

	// Create tables if they don't exist
	models := []interface{}{
		(*model.User)(nil),
		(*model.Thing)(nil),
	}

	for _, model := range models {
		if _, err := db.NewCreateTable().Model(model).IfNotExists().Exec(ctx); err != nil {
			return fmt.Errorf("failed to create table for model %T: %w", model, err)
		}
	}

	log.Println("Tables created successfully")
	return nil
}
