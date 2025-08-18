package service

import (
	"context"
	"fmt"
	"http_server/models"
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
func (s *userRepoImplStruct) UserRegister(ctx context.Context, user model.User) (*model.User, error) {
	if user.Role == "" || user.Name == "" || user.Email == "" {
		s.logger.ErrorContext(ctx, "User validation failed", user.Role, user.Name, user.Email, user)
		return nil, fmt.Errorf("user validation failed: missing fields:%s,%s,%s", user.Role, user.Email, user.Name)
	}
	user.Role = string(models.AdminRole)
	resp, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		s.logger.Error("failed to create user ", err)
		return nil, fmt.Errorf("failed to create user %w", err)

	}
	return resp, nil
}
