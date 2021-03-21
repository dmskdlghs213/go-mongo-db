package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoAccessor interface {
	Create(ctx context.Context, req *ReqCreate) error
	Creates(ctx context.Context, req []ReqCreate) error
	Find(ctx context.Context, name string) (*User, error)
	Finds(ctx context.Context, name ReqFinds) ([]User, error)
	Update(ctx context.Context, req *ReqUpdate) error
	Updates(ctx context.Context, req *ReqUpdates) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
