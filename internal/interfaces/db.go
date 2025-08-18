package interfaces

import (
	"context"

	"github.com/rafly-ananda/snappsy-uploader-api/internal/models"
)

type ImageRepository interface {
	Insert(ctx context.Context, image models.Images) (string, error)
	FindAllBySession(ctx context.Context, sessionId string) ([]models.Images, error)
	ObjSectStorage
}
