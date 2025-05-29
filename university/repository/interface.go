package repository

import (
	"context"
	"university/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UniversityRepository interface {
	CreateUniversity(ctx context.Context, university models.University) (*models.University, error)
	ListUniversity(ctx context.Context, pagination models.Pagination) ([]models.University, int64, error)
	UpdateUniversity(ctx context.Context, university models.University) (*models.University, error)
	GetUniversityById(ctx context.Context, id primitive.ObjectID) (*models.University, error)
	GetUniversityByEmail(ctx context.Context, email string) (*models.University, error)
	GetUniversityByName(ctx context.Context, name string) (*models.University, error)
}
