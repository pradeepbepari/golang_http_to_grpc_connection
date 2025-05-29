package database

import (
	"context"
	"database/sql"
	"log"
)

func Schemachange(ctx context.Context, db *sql.DB) {
	createUsersTable := `
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    contact VARCHAR(20),
    address TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);`

	if _, err := db.ExecContext(ctx, createUsersTable); err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}

}
