package models

type Post struct {
    Username string `form:"username" json:"username" binding:"required"`
    PhotoURL string `form:"photo_url" json:"photo_url" binding:"required"`
    Caption  string `form:"caption" json:"caption"`
}
