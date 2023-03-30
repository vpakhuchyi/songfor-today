package models

type Artist struct {
	ID            int    `json:"id"`
	Link          string `json:"link"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	PictureBig    string `json:"picture_big"`
	PictureMedium string `json:"picture_medium"`
	PictureSmall  string `json:"picture_small"`
	PictureXl     string `json:"picture_xl"`
	Radio         bool   `json:"radio"`
	Share         string `json:"share"`
	Tracklist     string `json:"tracklist"`
	Type          string `json:"type"`
}
