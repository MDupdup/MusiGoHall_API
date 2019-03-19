package models

type ArtistFull struct {
	Profile     string `json:"profile"`
	ReleasesURL string `json:"releases_url"`
	Name        string `json:"name"`
	URI         string `json:"uri"`
	Members     []struct {
		Active      bool   `json:"active"`
		ResourceURL string `json:"resource_url"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
	} `json:"members"`
	Urls   []string `json:"urls"`
	Images []struct {
		URI         string `json:"uri"`
		Height      int    `json:"height"`
		Width       int    `json:"width"`
		ResourceURL string `json:"resource_url"`
		Type        string `json:"type"`
		URI150      string `json:"uri150"`
	} `json:"images"`
	ResourceURL string `json:"resource_url"`
	ID          int    `json:"id"`
	DataQuality string `json:"data_quality"`
	Realname    string `json:"realname"`
}

type Artist struct {
	Thumb       string      `json:"thumb"`
	Title       string      `json:"title"`
	URI         string      `json:"uri"`
	MasterURL   interface{} `json:"master_url"`
	CoverImage  string      `json:"cover_image"`
	ResourceURL string      `json:"resource_url"`
	MasterID    interface{} `json:"master_id"`
	Type        string      `json:"type"`
	ID          int         `json:"id"`
}

type Release struct {
	Style      []string `json:"style"`
	MasterID   int      `json:"master_id"`
	Thumb      string   `json:"thumb"`
	Format     []string `json:"format"`
	Country    string   `json:"country"`
	Barcode    []string `json:"barcode"`
	URI        string   `json:"uri"`
	MasterURL  string   `json:"master_url"`
	Label      []string `json:"label"`
	CoverImage string   `json:"cover_image"`
	Catno      string   `json:"catno"`
	Community  struct {
		Have int `json:"have"`
		Want int `json:"want"`
	} `json:"community"`
	Year        string   `json:"year"`
	Genre       []string `json:"genre"`
	Title       string   `json:"title"`
	ResourceURL string   `json:"resource_url"`
	Type        string   `json:"type"`
	ID          int      `json:"id"`
}

type Label struct {
	Thumb       string      `json:"thumb"`
	Title       string      `json:"title"`
	URI         string      `json:"uri"`
	MasterURL   interface{} `json:"master_url"`
	CoverImage  string      `json:"cover_image"`
	ResourceURL string      `json:"resource_url"`
	MasterID    interface{} `json:"master_id"`
	Type        string      `json:"type"`
	ID          int         `json:"id"`
}

type Result struct {
	Pagination struct {
		PerPage int `json:"per_page"`
		Pages   int `json:"pages"`
		Page    int `json:"page"`
		Urls    struct {
			Last string `json:"last"`
			Next string `json:"next"`
		} `json:"urls"`
		Items int `json:"items"`
	} `json:"pagination"`
	Results []struct {
		Thumb       string      `json:"thumb"`
		Title       string      `json:"title"`
		URI         string      `json:"uri"`
		MasterURL   interface{} `json:"master_url"`
		CoverImage  string      `json:"cover_image"`
		ResourceURL string      `json:"resource_url"`
		MasterID    interface{} `json:"master_id"`
		Type        string      `json:"type"`
		ID          int         `json:"id"`
	} `json:"results"`
}
