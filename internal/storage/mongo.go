package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func NewMongo(host string, username string, password string, dbName string, mongoOpts string) (*MongoStorage, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s?%s", username, password, host, dbName, mongoOpts)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Mongo connection error:", err)
		return nil, err
	}

	err = cl.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Mongo ping error:", err)
		return nil, err
	}

	log.Println("Connected to MongoDB:", uri)

	db := cl.Database(dbName)

	return &MongoStorage{
		Client: cl,
		Db:     db,
	}, nil
}
