package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rafly-ananda/snappsy-uploader-api/internal/config"
	"github.com/rafly-ananda/snappsy-uploader-api/internal/http/handlers/images"
	"github.com/rafly-ananda/snappsy-uploader-api/internal/repositories"
	"github.com/rafly-ananda/snappsy-uploader-api/internal/services"
	"github.com/rafly-ananda/snappsy-uploader-api/internal/storage"
)

func NewRouter(minio *storage.MinioStorage, mongo *storage.MongoStorage, cfg *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Repository Initialization
	imageRepo := repositories.NewImageRepository(mongo, minio, cfg.MongoCfg)

	// Service Initialization
	imageSvc := services.NewImageService(imageRepo)

	// Handler Initialization
	imageHandler := images.NewImageHandler(imageSvc)

	api := r.Group("/api")

	{
		// Images
		api.POST("/images", imageHandler.CommitImageUpload)
		api.POST("/images/presigned-uploader", imageHandler.GeneratePresignedUploader)
		api.GET("/generate", imageHandler.GeneratePresignedViewer)
		api.GET("/images", imageHandler.GetAllImages)

	}

	// simple health check
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, gin.H{
			"status": "OK",
		})
	})

	return r
}
