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
	Title      string `json:"title"`
	URI        string `json:"uri"`
	CoverImage string `json:"cover_image"`
	ID         int    `json:"id"`
}

type Release struct {
	Status string `json:"status"`
	Videos []struct {
		Duration    int    `json:"duration"`
		Embed       bool   `json:"embed"`
		Title       string `json:"title"`
		Description string `json:"description"`
		URI         string `json:"uri"`
	} `json:"videos"`
	Series []interface{} `json:"series"`
	Labels []struct {
		Name           string `json:"name"`
		EntityType     string `json:"entity_type"`
		Catno          string `json:"catno"`
		ResourceURL    string `json:"resource_url"`
		ID             int    `json:"id"`
		EntityTypeName string `json:"entity_type_name"`
	} `json:"labels"`
	Year      int `json:"year"`
	Community struct {
		Status string `json:"status"`
		Rating struct {
			Count   int     `json:"count"`
			Average float64 `json:"average"`
		} `json:"rating"`
		Want         int `json:"want"`
		Contributors []struct {
			Username    string `json:"username"`
			ResourceURL string `json:"resource_url"`
		} `json:"contributors"`
		Have      int `json:"have"`
		Submitter struct {
			Username    string `json:"username"`
			ResourceURL string `json:"resource_url"`
		} `json:"submitter"`
		DataQuality string `json:"data_quality"`
	} `json:"community"`
	Artists []struct {
		Join        string `json:"join"`
		Name        string `json:"name"`
		Anv         string `json:"anv"`
		Tracks      string `json:"tracks"`
		Role        string `json:"role"`
		ResourceURL string `json:"resource_url"`
		ID          int    `json:"id"`
	} `json:"artists"`
	Images []struct {
		URI         string `json:"uri"`
		Height      int    `json:"height"`
		Width       int    `json:"width"`
		ResourceURL string `json:"resource_url"`
		Type        string `json:"type"`
		URI150      string `json:"uri150"`
	} `json:"images"`
	FormatQuantity    int      `json:"format_quantity"`
	ID                int      `json:"id"`
	ArtistsSort       string   `json:"artists_sort"`
	Genres            []string `json:"genres"`
	Thumb             string   `json:"thumb"`
	NumForSale        int      `json:"num_for_sale"`
	Title             string   `json:"title"`
	DateChanged       string   `json:"date_changed"`
	MasterID          int      `json:"master_id"`
	LowestPrice       float64  `json:"lowest_price"`
	Styles            []string `json:"styles"`
	ReleasedFormatted string   `json:"released_formatted"`
	Formats           []struct {
		Qty          string   `json:"qty"`
		Descriptions []string `json:"descriptions"`
		Name         string   `json:"name"`
	} `json:"formats"`
	EstimatedWeight int    `json:"estimated_weight"`
	MasterURL       string `json:"master_url"`
	Released        string `json:"released"`
	DateAdded       string `json:"date_added"`
	Extraartists    []struct {
		Join        string `json:"join"`
		Name        string `json:"name"`
		Anv         string `json:"anv"`
		Tracks      string `json:"tracks"`
		Role        string `json:"role"`
		ResourceURL string `json:"resource_url"`
		ID          int    `json:"id"`
	} `json:"extraartists"`
	Tracklist []struct {
		Duration     string `json:"duration"`
		Position     string `json:"position"`
		Type         string `json:"type_"`
		Extraartists []struct {
			Join        string `json:"join"`
			Name        string `json:"name"`
			Anv         string `json:"anv"`
			Tracks      string `json:"tracks"`
			Role        string `json:"role"`
			ResourceURL string `json:"resource_url"`
			ID          int    `json:"id"`
		} `json:"extraartists"`
		Title string `json:"title"`
	} `json:"tracklist"`
	Notes       string `json:"notes"`
	Identifiers []struct {
		Type        string `json:"type"`
		Description string `json:"description"`
		Value       string `json:"value"`
	} `json:"identifiers"`
	Companies []struct {
		Name           string `json:"name"`
		EntityType     string `json:"entity_type"`
		Catno          string `json:"catno"`
		ResourceURL    string `json:"resource_url"`
		ID             int    `json:"id"`
		EntityTypeName string `json:"entity_type_name"`
	} `json:"companies"`
	URI         string `json:"uri"`
	Country     string `json:"country"`
	ResourceURL string `json:"resource_url"`
	DataQuality string `json:"data_quality"`
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
