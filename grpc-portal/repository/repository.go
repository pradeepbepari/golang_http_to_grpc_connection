package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const USERS_COLLECTION = "users"

type Repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		db: db,
	}
}
