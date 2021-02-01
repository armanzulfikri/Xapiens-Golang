package models

// photos model

type Photos struct {
	AlbumID      uint   `json:"albumID"`
	ID           uint   `json:"id"`
	Title        string `json:title`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl`
}
