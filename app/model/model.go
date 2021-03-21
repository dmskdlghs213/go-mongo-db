package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name"`
	Old  uint32             `json:"old"`
}
