package service

import (
	"context"
	"fmt"
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
func (s authService) LoginUser(ctx context.Context, email string) (models.User, error) {
	resp, err := s.repo.Login(ctx, email)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to login user with this ", email)
		return models.User{}, fmt.Errorf("failed to logi user with %s  and error %v", email, err)
	}
	return resp, nil

}
