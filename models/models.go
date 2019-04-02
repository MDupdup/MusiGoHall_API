package models

type Artist struct {
	Profile     string `json:"profile"`
	Name        string `json:"name"`
	Members     []Member
	ReleasesMin []ReleaseMin
	ID          int `json:"id"`
}

type Member struct {
	Active      bool   `json:"active"`
	ResourceURL string `json:"resource_url"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
}

type ReleaseMin struct {
	Title       string
	Thumb       string
	Year        int
	ResourceURL string
	ID          int
}

type Release struct {
	Title     string
	Thumb     string
	Year      int
	DateAdded string
	Country   string
	Style     []string
	TrackList []Track
	ID        int
}

type Track struct {
	Title    string
	Type     string
	Duration string
	Position string
}

type Result struct {
	Title      string `json:"title"`
	URI        string `json:"uri"`
	CoverImage string `json:"cover_image"`
	ID         int    `json:"id"`
}
