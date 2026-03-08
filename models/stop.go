package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transit struct {
	Mode  string `bson:"mode" json:"mode"`
	Time  string `bson:"time" json:"time"`
	Color string `bson:"color,omitempty" json:"color,omitempty"`
}

type Stop struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TripID      primitive.ObjectID `bson:"trip_id" json:"tripId"`
	UserID      primitive.ObjectID `bson:"user_id" json:"userId"`
	OrderIndex  int                `bson:"orderIndex" json:"orderIndex"`
	Day         int                `bson:"day" json:"day"`
	Time        string             `bson:"time" json:"time"`
	Name        string             `bson:"name" json:"name"`
	Status      string             `bson:"status" json:"status"`
	City        string             `bson:"city,omitempty" json:"city,omitempty"`
	Country     string             `bson:"country,omitempty" json:"country,omitempty"`
	Category    string             `bson:"category,omitempty" json:"category,omitempty"`
	IsCompleted bool               `bson:"isCompleted" json:"isCompleted"`
	IsActive    bool               `bson:"isActive" json:"isActive"`
	Image       string             `bson:"image,omitempty" json:"image,omitempty"`
	Actions     []string           `bson:"actions,omitempty" json:"actions,omitempty"`
	Transit     *Transit           `bson:"transit,omitempty" json:"transit,omitempty"`
	Lat         float64            `bson:"lat,omitempty" json:"lat,omitempty"`
	Lng         float64            `bson:"lng,omitempty" json:"lng,omitempty"`
}
