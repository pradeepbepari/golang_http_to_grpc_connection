package repository

import (
	"context"
	"grpc-portal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		db: db,
	}
}
func (r *Repository) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	user.Id = primitive.NewObjectID()
	collection := r.db.Collection("users")
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return &model.User{Id: user.Id}, nil

}
