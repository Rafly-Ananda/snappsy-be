package images

type CommitUploadReq struct {
	SessionId string `json:"sessionId" binding:"required"`
	Username  string `json:"username" binding:"required"`
	MinioKey  string `json:"minioKey" binding:"required"`
	Captions  string `json:"captions" binding:"required"`
}

type CommitUploadRes struct {
	ID string `json:"id"`
}
