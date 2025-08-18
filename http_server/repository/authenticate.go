package repository

import (
	"context"
	"database/sql"
	"fmt"
	"http_server/models"
)

func (r *repo) Login(ctx context.Context, email string) (models.User, error) {
	var userData models.User
	err := r.db.QueryRowContext(ctx, selectUser, email).Scan(
		&userData.ID,
		&userData.Name,
		&userData.Email,
		&userData.Password,
		&userData.Contact,
		&userData.Address,
		&userData.CreatedAt,
		&userData.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Errorf("No user found with email: %s", email)
			return userData, fmt.Errorf("userData not found with email: %s", email)
		}
		r.logger.Errorf("Database error while fetching userData with email %s: %v", email, err)
		return userData, fmt.Errorf("failed to fetch userData: %w", err)
	}

	r.logger.Infof("userData fetched successfully with email: %s", email)
	return userData, nil

}
