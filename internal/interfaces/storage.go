package interfaces

import (
	"context"
	"time"
)

type ObjectStorageUploader interface {
	Upload(ctx context.Context, key string, data []byte, contentType string) (string, error)
	PresignPut(ctx context.Context, key string, contentType string) (string, error)
}

type ObjectStorageViewer interface {
	PresignGet(ctx context.Context, key string, expiry time.Duration) (string, error)
}

// Compose, handler will receive this interface
type ObjectStorage interface {
	ObjectStorageUploader
	ObjectStorageViewer
}
