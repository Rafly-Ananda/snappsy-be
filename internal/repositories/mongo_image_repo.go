package repositories

import (
	"context"
	"time"

	"github.com/rafly-ananda/snappsy-uploader-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoImageRepository struct {
	collection *mongo.Collection
}

func NewMongoImageRepository(col *mongo.Collection) *MongoImageRepository {
	return &MongoImageRepository{collection: col}
}

func (r *MongoImageRepository) Insert(ctx context.Context, image models.Images) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := r.collection.InsertOne(ctx, image)
	if err != nil {
		return "", err
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return "", nil
}

func (r *MongoImageRepository) FindAllBySession(ctx context.Context, sessionId string) ([]models.Images, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cur, err := r.collection.Find(ctx, bson.M{"sessionId": sessionId})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var out []models.Images
	if err := cur.All(ctx, &out); err != nil {
		return nil, err
	}
	return out, nil
}
