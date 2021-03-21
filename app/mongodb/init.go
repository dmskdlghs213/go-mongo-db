package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Database

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	d := time.Duration(time.Second * 2)
	direct := true

	ops := &options.ClientOptions{
		ConnectTimeout: &d,
		Direct:         &direct,
	}

	client, err := mongo.Connect(ctx, ops, options.Client().ApplyURI("mongodb://root:root@mongo-server:27017"))
	if err != nil {
		log.Fatalf("connection error, err_msg:%v", err)
	}
	log.Println("connection success")

	conf := &mongodb.Config{
		DatabaseName: "test_dev",
	}

	driver, err := mongodb.WithInstance(client, conf)
	if err != nil {
		log.Fatalf("get instance error from mongodb client, err_msg:%v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "mongodb", driver)
	if err != nil {
		log.Fatalf("get instance error from mongodb driver, err_msg:%v", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalf("migration exec error, err_msg: %v", err)
	}

	Mongo = client.Database("test_dev")
}
