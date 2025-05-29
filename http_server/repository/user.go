package repository

import (
	"context"
	"fmt"
	"http_server/models"
)

func (r *repo) CreateUser(ctx context.Context, user models.User) error {

	query := `
        INSERT INTO users (id, name, email, password, contact, address, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW());
    `
	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.Name, user.Email, user.Password, user.Contact, user.Address,
	)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}
