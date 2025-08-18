package repository

import (
	"context"
	"fmt"
	"http_server/models"
)

func (r *repo) CreateUser(ctx context.Context, user models.User) error {

	_, err := r.db.ExecContext(ctx, createUser,
		user.ID, user.Name, user.Email, user.Password, user.Country, user.State, user.Role, user.Contact, user.Address,
	)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}
