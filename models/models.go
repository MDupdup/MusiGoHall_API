package models

type Image struct {
	Url  string
	Size string
}

type Track struct {
	Name     string
	Url      string
	Duration string
	Position string
}

type Artist struct {
	Name      string
	Mbid      string
	Url       string
	PlayCount string
	Images    []Image
	IsOnTour  bool
	Similar   []Artist
	Summup    string
	Content   string
	Albums    []Album
	Tags      []string
}

type Album struct {
	Name    string
	Mbid    string
	Artist  string
	Url     string
	Images  []Image
	Tracks  []Track
	Summup  string
	Content string
	Tags    []string
}
