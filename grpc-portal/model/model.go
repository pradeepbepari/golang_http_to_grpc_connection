package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Country   string             `bson:"country"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Contact   string             `bson:"contact"`
	Address   string             `bson:"address"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}
