package repository

import (
	"context"
	"time"
	"university/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	UNIVERSITY_COLLECTION = "University"
)

func (r *repo) CreateUniversity(ctx context.Context, university models.University) (*models.University, error) {
	university.ID = primitive.NewObjectID()
	university.CreatedAt = time.Now()
	university.UpdatedAt = time.Now()
	_, err := r.db.Collection(UNIVERSITY_COLLECTION).InsertOne(ctx, university)
	if err != nil {
		return nil, err
	}
	return &university, nil
}
func (r *repo) ListUniversity(ctx context.Context, pagination models.Pagination) ([]models.University, int64, error) {
	var universities []models.University
	totalCount, err := r.db.Collection(UNIVERSITY_COLLECTION).CountDocuments(ctx, nil)
	if err != nil {
		return nil, 0, err
	}
	skip := (pagination.PageSize - 1) * pagination.PageSize
	limit := pagination.PageSize
	cursor, err := r.db.Collection(UNIVERSITY_COLLECTION).Find(ctx, nil, options.Find().SetLimit(int64(limit)).SetSkip(int64(skip)))
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var university models.University
		if err := cursor.Decode(&university); err != nil {
			return nil, 0, err
		}
		universities = append(universities, university)
	}
	if err := cursor.Err(); err != nil {
		return nil, 0, err
	}

	return universities, totalCount, nil
}
func (r *repo) UpdateUniversity(ctx context.Context, university models.University) (*models.University, error) {
	university.UpdatedAt = time.Now()
	filter := primitive.M{"_id": university.ID}
	update := primitive.M{
		"$set": university,
	}
	_, err := r.db.Collection(UNIVERSITY_COLLECTION).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return &university, nil
}
func (r *repo) GetUniversityById(ctx context.Context, id primitive.ObjectID) (*models.University, error) {
	var university models.University
	err := r.db.Collection(UNIVERSITY_COLLECTION).FindOne(ctx, primitive.M{"_id": id}).Decode(&university)
	if err != nil {
		return nil, err
	}
	return &university, nil
}
func (r *repo) GetUniversityByEmail(ctx context.Context, email string) (*models.University, error) {
	var university models.University
	err := r.db.Collection(UNIVERSITY_COLLECTION).FindOne(ctx, primitive.M{"email": email}).Decode(&university)
	if err != nil {
		return nil, err
	}
	return &university, nil
}

func (r *repo) GetUniversityByName(ctx context.Context, name string) (*models.University, error) {
	var university models.University
	err := r.db.Collection(UNIVERSITY_COLLECTION).FindOne(ctx, primitive.M{"name": primitive.Regex{Pattern: "^" + name + "$", Options: ""}}).Decode(&university)
	if err != nil {
		return nil, err
	}
	return &university, nil
}
