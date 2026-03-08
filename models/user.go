package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password,omitempty" json:"-"` // Never send password in JSON
	Name     string             `bson:"name" json:"name"`
	Picture  string             `bson:"picture" json:"picture"`
}
