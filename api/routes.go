package api

import (
	"MusiGoHall_API/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var baseUrl = "http://ws.audioscrobbler.com/2.0"
var key = "f6b194e8e973a19daa39f08ee677c5f0"
var secret = "b2204a511b7c2f8db6452ca4bb28dcda"

/**
Get given release by id
*/
//TODO make transition for Last.fm!
func GetAlbum(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	method := "album.getInfo"

	fmt.Println(params["type"], ":", params["name"])
	url := fmt.Sprintf("%s/?method=%s&mbid=%s&api_key=%s&format=json", baseUrl, method, params["id"], key)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest:", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	gresult := gjson.Parse(bodyString).Get("album")

	var trackList []models.Track
	for _, result := range gresult.Get("tracks").Get("track").Array() {
		trackList = append(trackList, models.Track{
			Name:     result.Get("name").Str,
			Url:      result.Get("url").Str,
			Duration: result.Get("duration").Str,
			Position: result.Get("@attr").Get("rank").Str,
		})
	}

	var tags []models.Tag
	for _, result := range gresult.Get("tags").Get("tag").Array() {
		tags = append(tags, models.Tag{
			Name: result.Get("name").Str,
			Url:  result.Get("url").Str,
		})
	}

	var images []models.Image
	for _, result := range gresult.Get("image").Array() {
		images = append(images, models.Image{
			Url:  result.Get("#text").Str,
			Size: result.Get("size").Str,
		})
	}

	var album = models.Album{
		Name:    gresult.Get("name").Str,
		Mbid:    gresult.Get("mbid").Str,
		Artist:  gresult.Get("artist").Str,
		Url:     gresult.Get("url").Str,
		Images:  images,
		Tracks:  trackList,
		Summup:  gresult.Get("wiki").Get("summary").Str,
		Content: gresult.Get("wiki").Get("content").Str,
		Tags:    tags,
	}

	err = json.NewEncoder(w).Encode(album)
	if err != nil {
		log.Fatal("jsonEncode:", err)
	}
}

/**
Get given artist by id
*/
func GetArtist(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	method := "artist.getInfo"

	fmt.Println(params["type"], ":", params["name"])
	url := fmt.Sprintf("%s/?method=%s&mbid=%s&api_key=%s&format=json", baseUrl, method, params["id"], key)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest:", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	gresult := gjson.Parse(bodyString).Get("artist")

	var similarArtists []models.Artist
	for _, result := range gresult.Get("similar").Get("artist").Array() {
		var images []models.Image
		for _, iresult := range result.Get("image").Array() {
			images = append(images, models.Image{
				Url:  iresult.Get("#text").Str,
				Size: iresult.Get("size").Str,
			})
		}

		similarArtists = append(similarArtists, models.Artist{
			Name:   result.Get("name").Str,
			Url:    result.Get("url").Str,
			Images: images,
		})
	}

	var images []models.Image
	for _, iresult := range gresult.Get("image").Array() {
		images = append(images, models.Image{
			Url:  iresult.Get("#text").Str,
			Size: iresult.Get("size").Str,
		})
	}
	var artist = models.Artist{
		Name:      gresult.Get("name").Str,
		Mbid:      gresult.Get("mbid").Str,
		Url:       gresult.Get("url").Str,
		PlayCount: gresult.Get("stats").Get("playcount").Str,
		Images:    images,
		IsOnTour:  gresult.Get("ontour").Bool(),
		Similar:   similarArtists,
		Summup:    gresult.Get("summary").Str,
		Content:   gresult.Get("content").Str,
	}

	err = json.NewEncoder(w).Encode(artist)
	if err != nil {
		log.Fatal("jsonEncode:", err)
	}
}

/*func getArtistReleases(url string) (res []models.ReleaseMin) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest:", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)

	gresult := gjson.Get(bodyString, "releases")
	for _, result := range gresult.Array() {

		if result.Get("role").Str == "Main" {
			res = append(res, models.ReleaseMin{
				Title:       result.Get("title").Str,
				Thumb:       result.Get("thumb").Str,
				Year:        int(result.Get("year").Num),
				ResourceURL: result.Get("resource_url").Str,
				ID:          int(result.Get("id").Num),
			})
		}
	}

	return res
}*/

/**
Search for releases by value
*/
func SearchAlbum(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	method := "album.search"

	url := fmt.Sprintf("%s/?method=%s&album=%s&api_key=%s&format=json", baseUrl, method, params["value"], key)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest:", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)

	var albums []models.Album
	gresult := gjson.Get(bodyString, "results").Get("albummatches").Get("album")
	for _, result := range gresult.Array() {

		var images []models.Image
		for _, iresult := range result.Get("style").Array() {
			images = append(images, models.Image{
				Url:  iresult.Get("#text").Str,
				Size: iresult.Get("size").Str,
			})
		}

		albums = append(albums, models.Album{
			Name:   result.Get("name").Str,
			Mbid:   result.Get("mbid").Str,
			Artist: result.Get("artist").Str,
			Url:    result.Get("url").Str,
			Images: images,
		})
	}

	err = json.NewEncoder(w).Encode(albums)
	if err != nil {
		log.Fatal("jsonEncode:", err)
		return
	}
}

/**
Search for artists by value
*/
func SearchArtist(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	parameter, err := url.Parse(url.QueryEscape(params["value"]))
	if err != nil {
		log.Fatal("ParseError:", err)
		panic(err)
	}
	fmt.Printf("get artist: [%s]", parameter)

	method := "artist.search"

	url := fmt.Sprintf("%s/?method=%s&artist=%s&api_key=%s&format=json", baseUrl, method, parameter.String(), key)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest:", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)

	//var regex = regexp.MustCompile(`^.*?\([^\d]*(\d+)[^\d]*\).*$`)

	var artists []models.Artist
	gresult := gjson.Get(bodyString, "results").Get("artistmatches").Get("artist")
	for _, result := range gresult.Array() {

		var images []models.Image
		for _, iresult := range result.Get("image").Array() {
			images = append(images, models.Image{
				Url:  iresult.Get("#text").Str,
				Size: iresult.Get("size").Str,
			})
		}

		artists = append(artists, models.Artist{
			Name:   result.Get("name").Str,
			Mbid:   result.Get("mbid").Str,
			Url:    result.Get("url").Str,
			Images: images,
		})
	}

	err = json.NewEncoder(w).Encode(artists)
	if err != nil {
		log.Fatal("jsonEncode:", err)
		return
	}
}
