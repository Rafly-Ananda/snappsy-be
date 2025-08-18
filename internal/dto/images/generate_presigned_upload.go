package images

type GeneratePresignedUrlReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	PhotoURL string `form:"photoUrl" json:"photoUrl" binding:"required"`
	EventKey string `form:"eventKey" json:"eventKey" binding:"required"`
}

type GeneratePresignedUrlRes struct {
	UploadUrl   string
	ObjectKey   string
	ContentType string
}
