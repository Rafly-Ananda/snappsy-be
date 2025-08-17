package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rafly-ananda/snappsy-uploader-api/internal/http/handlers"
	"github.com/rafly-ananda/snappsy-uploader-api/internal/interfaces"
)

func NewRouter(u interfaces.ObjectStorage) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	objectStorageUploader := handlers.NewPostHandler(u)
	objectStorageGenerator := handlers.NewGenerateHandler(u)

	api := r.Group("/api")

	{
		api.POST("/post", objectStorageUploader.CreatePost)
		api.GET("/generate", objectStorageGenerator.GeneratePresignedUrl)
	}

	// simple health check
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, gin.H{
			"status": "OK",
		})
	})

	return r
}
