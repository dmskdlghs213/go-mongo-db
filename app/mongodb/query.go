package mongodb

import (
	"context"
	"fmt"

	"github.com/dmskdlghs213/go-mongoDB/app/model"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoQuery struct {
	mongo *mongo.Database
}

func NewMongoQuery(m *mongo.Database) *MongoQuery {
	return &MongoQuery{
		mongo: m,
	}
}

func (m *MongoQuery) Create(ctx context.Context, req *model.ReqCreate) error {

	req.ID = primitive.NewObjectID()

	res, err := m.mongo.Collection("test_collections").InsertOne(ctx, req)
	if err != nil {
		log.Error(err)
		return err
	}

	id := res.InsertedID
	fmt.Println(id)

	return nil
}

func (m *MongoQuery) Creates(ctx context.Context, req []model.ReqCreate) error {

	list := make([]interface{}, 0, 10)
	for i := range req {
		list = append(list, req[i])
	}

	res, err := m.mongo.Collection("test_collections").InsertMany(ctx, list)
	if err != nil {
		log.Error(err)
		return err
	}

	id := res.InsertedIDs
	fmt.Println(id)

	return nil
}

func (m *MongoQuery) Find(ctx context.Context, name string) (*model.User, error) {

	filter := bson.D{{Key: "name", Value: name}}
	result := m.mongo.Collection("test_collections").FindOne(ctx, filter)
	if result.Err() != nil {
		log.Error(result.Err())
		return nil, result.Err()
	}

	var u model.User
	if err := result.Decode(&u); err != nil {
		log.Error(err)
		return nil, err
	}

	return &u, nil
}

func (m *MongoQuery) Finds(ctx context.Context, req model.ReqFinds) ([]model.User, error) {

	names := make([]string, 0, len(req.Names))
	for i := range req.Names {
		name := req.Names[i]
		names = append(names, name)
	}
	filter := bson.M{"name": bson.M{"$in": names}}

	cur, err := m.mongo.Collection("test_collections").Find(ctx, filter)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cur.Close(ctx)

	users := make([]model.User, 0, len(req.Names))
	for cur.Next(ctx) {
		var u model.User
		if err := cur.Decode(&u); err != nil {
			log.Error(err)
			return nil, err
		}

		users = append(users, u)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *MongoQuery) Update(ctx context.Context, req *model.ReqUpdate) error {

	// filter := bson.D{bson.E{Key: "name", Value: bson.M{"$eq": req.Name}}}
	filter := bson.D{bson.E{Key: "_id", Value: req.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "address", Value: "埼玉県"},
		primitive.E{Key: "old", Value: req.Old},
	}}}

	_, err := m.mongo.Collection("test_collections").UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (m *MongoQuery) Updates(ctx context.Context, req *model.ReqUpdates) error {

	filter := bson.D{bson.E{Key: "_id", Value: bson.M{"$in": req.ID}}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "address", Value: "長野県"},
		primitive.E{Key: "old", Value: req.Old},
	}}}

	_, err := m.mongo.Collection("test_collections").UpdateMany(ctx, filter, update)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (m *MongoQuery) Delete(ctx context.Context, id primitive.ObjectID) error {

	filter := bson.D{bson.E{Key: "_id", Value: id}}
	_, err := m.mongo.Collection("test_collections").DeleteOne(ctx, filter)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
