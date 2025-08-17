package storage

import (
	"bytes"
	"context"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioStorage struct {
	client *minio.Client
	bucket string
}

func NewMinio(endpoint, accessKey, secretKey, bucket string, minioExpiry int, secure bool) (*MinioStorage, error) {
	cl, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: secure,
	})

	if err != nil {
		return nil, err
	}

	return &MinioStorage{client: cl, bucket: bucket}, nil
}

func (m *MinioStorage) Upload(ctx context.Context, key string, data []byte, contentType string) (string, error) {
	_, err := m.client.PutObject(ctx, m.bucket, key, bytes.NewReader(data), int64(len(data)),
		minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		return "", err
	}

	return key, nil
}

func (m *MinioStorage) PresignPut(ctx context.Context, key, contentType string) (string, error) {
	url, err := m.client.PresignedPutObject(ctx, m.bucket, key, 10*time.Minute)

	if err != nil {
		return "", err
	}

	return url.String(), nil
}

func (m *MinioStorage) PresignGet(ctx context.Context, key string, expiry time.Duration) (string, error) {
	u, err := m.client.PresignedGetObject(ctx, m.bucket, key, expiry, nil)

	if err != nil {
		return "", err
	}

	return u.String(), nil
}
