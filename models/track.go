package models

type Track struct {
	ID                    int           `json:"id"`
	Album                 Album         `json:"album"`
	Artist                Artist        `json:"artist"`
	AvailableCountries    []string      `json:"available_countries"`
	BPM                   float64       `json:"bpm"`
	Contributors          []Contributor `json:"contributors"`
	DiskNumber            int           `json:"disk_number"`
	Duration              int           `json:"duration"`
	ExplicitContentCover  int           `json:"explicit_content_cover"`
	ExplicitContentLyrics int           `json:"explicit_content_lyrics"`
	ExplicitLyrics        bool          `json:"explicit_lyrics"`
	Gain                  float64       `json:"gain"`
	ISRC                  string        `json:"isrc"`
	Link                  string        `json:"link"`
	MD5Image              string        `json:"md5_image"`
	Preview               string        `json:"preview"`
	Rank                  int           `json:"rank"`
	Readable              bool          `json:"readable"`
	ReleaseDate           string        `json:"release_date"`
	Share                 string        `json:"share"`
	Title                 string        `json:"title"`
	TitleShort            string        `json:"title_short"`
	TitleVersion          string        `json:"title_version"`
	TrackPosition         int           `json:"track_position"`
	Type                  string        `json:"type"`
}
