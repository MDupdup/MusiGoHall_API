package models

type Image struct {
	Url  string
	Size string
}

type Track struct {
	Name     string
	Url      string
	Duration int
	Position string
}

type Artist struct {
	Name      string   `bson:"Name"`
	Mbid      string   `bson:"Mbid"`
	Url       string   `bson:"Url"`
	PlayCount string   `bson:"PlayCount"`
	Images    []Image  `bson:"Images"`
	IsOnTour  bool     `bson:"IsOnTour"`
	Similar   []Artist `bson:"similar"`
	Summup    string   `bson:"Summup"`
	Content   string   `bson:"Content"`
	Albums    []Album  `bson:"Albums"`
	Tags      []string `bson:"Tags"`
}

type Album struct {
	Name    string   `bson:"Name"`
	Mbid    string   `bson:"Mbid"`
	Artist  string   `bson:"Artist"`
	Url     string   `bson:"Url"`
	Images  []Image  `bson:"Images"`
	Tracks  []Track  `bson:"Tracks"`
	Summup  string   `bson:"Summup"`
	Content string   `bson:"Content"`
	Tags    []string `bson:"Tags"`
}
