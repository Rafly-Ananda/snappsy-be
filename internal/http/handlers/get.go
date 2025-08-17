package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafly-ananda/snappsy-uploader-api/internal/interfaces"
)

type ObjectGetHandler struct {
	generator interfaces.ObjectStorageViewer
}

func NewGenerateHandler(u interfaces.ObjectStorageViewer) *ObjectGetHandler {
	return &ObjectGetHandler{generator: u}
}

func (h *ObjectGetHandler) GeneratePresignedUrl(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		c.JSON(400, gin.H{"error": "missing key"})
		return
	}

	presignedUrl, err := h.generator.PresignGet(c, key, 10*time.Minute)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"upload_url": presignedUrl,
	})
}
