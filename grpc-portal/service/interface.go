package service

import (
	"context"
	"grpc-portal/model"
)

type UserService interface {
	UserRegister(c context.Context, user model.User) (*model.User, error)
}
