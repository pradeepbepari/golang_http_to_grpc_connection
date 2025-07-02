package database

import (
	"context"
	"database/sql"
	"log"
)

func Schemachange(ctx context.Context, db *sql.DB) {
	if _, err := db.ExecContext(ctx, createUsersTable); err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}
}
