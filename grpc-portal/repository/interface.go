package repository

import (
	"context"
	"grpc-portal/model"
)

type UserRepository interface {
	CreateUser(c context.Context, user model.User) (*model.User, error)
}
