package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Trip struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID primitive.ObjectID `bson:"user_id" json:"userId"`
	Title  string             `bson:"title" json:"title"`
	Status string             `bson:"status" json:"status"`
    IsPublic bool             `bson:"is_public" json:"isPublic"`
}
