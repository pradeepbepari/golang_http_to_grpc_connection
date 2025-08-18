package repository

import (
	"context"
	"grpc-portal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	user.Id = primitive.NewObjectID()
	collection := r.db.Collection(USERS_COLLECTION)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return &model.User{Id: user.Id, Name: user.Name, Email: user.Email}, nil
}
