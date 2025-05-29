package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type University struct {
	ID              primitive.ObjectID `bson:"_id"`
	Name            string             `bson:"name"`
	Email           string             `bson:"email"`
	Contact         string             `bson:"contact"`
	Password        string             `bson:"password"`
	Logo            string             `bson:"logo"`
	Location        string             `bson:"location"`
	City            string             `bson:"city"`
	Website         string             `bson:"website"`
	EstablishedYear int32              `bson:"established_year"`
	Colleges        []string           `bson:"colleges"`
	Programs        []string           `bson:"programs"`
	Ranking         int32              `bson:"ranking"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
}
type Pagination struct {
	PageSize  int32
	PageLimit int32
}
type ListUniversity struct {
	University University
	Pagination Pagination
	Total      int32
}
