package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rafly-ananda/snappsy-uploader-api/internal/http/handlers/images"
)

type Handlers struct {
	Images *images.ImageHandler
}

func NewRouter(h Handlers) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		// Images
		img := v1.Group("/images")

		img.POST("/generate-uploader-url", h.Images.GeneratePresignedUploader)
		img.POST("", h.Images.CommitImageUpload)
		img.GET("/generate-url", h.Images.GeneratePresignedViewer)
		img.GET("/images", h.Images.GetAllImages)
	}

	// simple health check
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	return r
}
