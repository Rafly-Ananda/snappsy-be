package interfaces

import (
	"context"
	"time"
)

type ObjectStorageUploader interface {
	PresignedPut(ctx context.Context, key string, contentType string) (string, error)
}

type ObjectStorageViewer interface {
	PresignedGet(ctx context.Context, key string, expiry time.Duration) (string, error)
}

type ObjSectStorage interface {
	ObjectStorageUploader
	ObjectStorageViewer
}
