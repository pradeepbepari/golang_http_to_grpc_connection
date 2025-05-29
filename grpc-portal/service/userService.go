package service

import (
	"context"
	"sdk-helper/logger"

	"grpc-portal/model"
	"grpc-portal/repository"
)

type userRepoImplStruct struct {
	userRepo repository.UserRepository
	logger   *logger.Logger
}

func NewUserService(userRepo repository.UserRepository, logger *logger.Logger) UserService {
	return &userRepoImplStruct{
		userRepo: userRepo,
		logger:   logger,
	}
}
func (s *userRepoImplStruct) UserRegister(c context.Context, user model.User) (*model.User, error) {
	return s.userRepo.CreateUser(c, user)
}
