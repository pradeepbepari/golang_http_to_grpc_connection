package service

import (
	"context"
	"http_server/models"
	"http_server/repository"
	"sdk-helper/logger"
)

type authService struct {
	repo   repository.Repository
	logger *logger.Logger
}

func NewAuthService(repo repository.Repository, logger *logger.Logger) AuthService {
	return authService{
		repo:   repo,
		logger: logger,
	}
}
func (s authService) Login(context.Context, models.User) (string, error) {

	return "token", nil
}
