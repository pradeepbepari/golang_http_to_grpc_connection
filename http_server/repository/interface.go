package repository

import (
	"context"
	"http_server/models"
)

type Repository interface {
	UserRepository
	FileRepository
	AuthRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
}
type FileRepository interface {
}
type AuthRepository interface {
	Login(ctx context.Context, email string) (models.User, error)
}
