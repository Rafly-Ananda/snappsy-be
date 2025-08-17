package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafly-ananda/snappsy-uploader-api/internal/interfaces"
	"github.com/rafly-ananda/snappsy-uploader-api/internal/models"
)

type ObjectPostHandler struct {
	uploader interfaces.ObjectStorageUploader
}

func NewPostHandler(u interfaces.ObjectStorageUploader) *ObjectPostHandler {
	return &ObjectPostHandler{uploader: u}
}

// CreatePost receives form data: photo_url, username, caption.
// Content-Types supported: application/x-www-form-urlencoded, multipart/form-data
func (h *ObjectPostHandler) CreatePost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Basic validation (non-fatal: just to help clients)
	if post.PhotoURL == "" || post.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "photo_url and username are required",
		})
		return
	}

	// Generate unique object key
	key := post.Username + "-" + time.Now().Format("20060102150405") + ".png"

	// Get presigned upload URL from storage (via the interface)
	url, err := h.uploader.PresignPut(context.Background(), key, "image/png")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"upload_url": url,
		"object_key": key,
	})
}
