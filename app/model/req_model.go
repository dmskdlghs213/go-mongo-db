package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReqCreate struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `json:"name" bson:"name"`
	Old  uint32             `json:"old" bson:"user_old"`
}

type ReqFinds struct {
	Names []string `query:"name"`
}

type ReqUpdate struct {
	ID   primitive.ObjectID
	Name string `json:"name" bson:"name"`
	Old  uint32 `json:"old" bson:"old"`
}

type ReqUpdates struct {
	ID    []primitive.ObjectID
	ReqID []string `json:"ids"`
	Name  string   `json:"name" bson:"name"`
	Old   uint32   `json:"old" bson:"old"`
}
