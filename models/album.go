package models

type Album struct {
	ID          int    `json:"id"`
	Cover       string `json:"cover"`
	CoverBig    string `json:"cover_big"`
	CoverMedium string `json:"cover_medium"`
	CoverSmall  string `json:"cover_small"`
	CoverXl     string `json:"cover_xl"`
	Link        string `json:"link"`
	MD5Image    string `json:"md5_image"`
	ReleaseDate string `json:"release_date"`
	Title       string `json:"title"`
	Tracklist   string `json:"tracklist"`
	Type        string `json:"type"`
}
