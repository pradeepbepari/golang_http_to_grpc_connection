package service

import (
	"context"
	"sdk-helper/logger"
	"university/models"
	"university/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type universityService struct {
	repo   repository.UniversityRepository
	logger *logger.Logger
}

func NewUniversityService(repo repository.UniversityRepository, logger *logger.Logger) UniversityService {
	return &universityService{
		repo:   repo,
		logger: logger,
	}
}
func (u universityService) CreateUniversity(ctx context.Context, university models.University) (*models.University, error) {
	createdUniversity, err := u.repo.CreateUniversity(ctx, university)
	if err != nil {
		u.logger.Error("Error creating university", "error", err)
		return nil, err
	}
	return createdUniversity, nil

}
func (u universityService) ListUniversity(ctx context.Context, pagination models.Pagination) ([]models.University, int64, error) {
	universities, count, err := u.repo.ListUniversity(ctx, pagination)
	if err != nil {
		u.logger.Error("Error listing universities", "error", err)
		return nil, 0, err
	}
	return universities, count, nil
}
func (u universityService) UpdateUniversity(ctx context.Context, university models.University) (*models.University, error) {
	updatedUniversity, err := u.repo.UpdateUniversity(ctx, university)
	if err != nil {
		u.logger.Error("Error updating university", "error", err)
		return nil, err
	}
	return updatedUniversity, nil
}
func (u universityService) GetUniversityById(ctx context.Context, id primitive.ObjectID) (*models.University, error) {
	updatedUniversity, err := u.repo.GetUniversityById(ctx, id)
	if err != nil {
		u.logger.Error("Error getting university by ID %s", "error %v", id, err)
		return nil, err
	}
	return updatedUniversity, nil
}
func (u universityService) GetUniversityByEmail(ctx context.Context, email string) (*models.University, error) {
	updatedUniversity, err := u.repo.GetUniversityByEmail(ctx, email)
	if err != nil {
		u.logger.Error("Error getiing university by email %s", "error", email, err)
		return nil, err
	}
	return updatedUniversity, nil
}
func (u universityService) GetUniversityByName(ctx context.Context, name string) (*models.University, error) {
	updatedUniversity, err := u.repo.GetUniversityByName(ctx, name)
	if err != nil {
		u.logger.Error("Error updating university", "error", err)
		return nil, err
	}
	return updatedUniversity, nil
}
