package repository

import (
	"context"
	"http_server/models"
)

type Repository interface {
	UserRepository
	FileRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
}
type FileRepository interface {
}
