package repository

import "go.mongodb.org/mongo-driver/mongo"

type repo struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) UniversityRepository {
	return &repo{
		db: db,
	}
}
