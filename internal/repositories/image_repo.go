package repositories

import (
	"context"
	"time"

	"github.com/rafly-ananda/snappsy-uploader-api/internal/config"
	"github.com/rafly-ananda/snappsy-uploader-api/internal/models"
	"github.com/rafly-ananda/snappsy-uploader-api/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ImageRepository struct {
	collection *mongo.Collection
	client     *storage.MinioStorage
}

func NewImageRepository(db *storage.MongoStorage, objs *storage.MinioStorage, cfg config.MongoConfig) *ImageRepository {
	return &ImageRepository{
		collection: db.Db.Collection(cfg.ImageCollection),
		client:     objs,
	}
}

func (r *ImageRepository) Insert(ctx context.Context, image models.Images) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := r.collection.InsertOne(ctx, image)
	if err != nil {
		return "", err
	}

	// Convert to string ID
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return "", nil
}

func (r *ImageRepository) FindAllBySession(ctx context.Context, sessionId string) ([]models.Images, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"sessionId": sessionId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var images []models.Images
	if err := cursor.All(ctx, &images); err != nil {
		return nil, err
	}

	return images, nil
}

func (r *ImageRepository) PresignedPut(ctx context.Context, key string, contentType string) (string, error) {
	url, err := r.client.Client.PresignedPutObject(ctx, r.client.Bucket, key, 10*time.Minute)

	if err != nil {
		return "", err
	}

	return url.String(), nil
}

func (r *ImageRepository) PresignedGet(ctx context.Context, key string, expiry time.Duration) (string, error) {
	u, err := r.client.Client.PresignedGetObject(ctx, r.client.Bucket, key, expiry, nil)

	if err != nil {
		return "", err
	}

	return u.String(), nil
}
