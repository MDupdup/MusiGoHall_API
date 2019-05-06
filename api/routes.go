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
	. "net/url"
	"regexp"
)

var baseUrl = "http://ws.audioscrobbler.com/2.0"
var key = "f6b194e8e973a19daa39f08ee677c5f0"
var secret = "b2204a511b7c2f8db6452ca4bb28dcda"

/**
Get given release by id
*/
//TODO make transition for Last.fm!
func GetRelease(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	method := "album.getInfo"

	fmt.Println(params["type"], ":", params["name"])
	url := fmt.Sprintf("%s/?method=%s/%s", baseUrl, method, params["id"])

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
	gresult := gjson.Parse(bodyString)

	var trackList []models.Track
	for _, result := range gresult.Get("tracklist").Array() {
		trackList = append(trackList, models.Track{
			Title:    result.Get("title").Str,
			Type:     result.Get("type_").Str,
			Duration: result.Get("duration").Str,
			Position: result.Get("position").Str,
		})
	}

	var styles []string
	for _, result := range gresult.Get("styles").Array() {
		styles = append(styles, result.Str)
	}

	var release = models.Release{
		Title:     gresult.Get("title").Str,
		Thumb:     gresult.Get("thumb").Str,
		Year:      int(gresult.Get("year").Num),
		DateAdded: gresult.Get("date_added").Str,
		Country:   gresult.Get("country").Str,
		Style:     styles,
		TrackList: trackList,
		ID:        int(gresult.Get("id").Num),
	}

	err = json.NewEncoder(w).Encode(release)
	if err != nil {
		log.Fatal("jsonEncode:", err)
	}
}

/**
Get given artist by id
*/
func GetArtist(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	fmt.Println(params["type"], ":", params["name"])
	url := fmt.Sprintf("%s/artists/%s", baseUrl, params["id"])

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
	gresult := gjson.Parse(bodyString)

	var members []models.Member
	for _, result := range gresult.Get("members").Array() {
		members = append(members, models.Member{
			result.Get("active").Bool(),
			result.Get("resource_url").Str,
			int(result.Get("Id").Num),
			result.Get("name").Str,
		})
	}

	var artist = models.Artist{
		gresult.Get("profile").Str,
		gresult.Get("name").Str,
		members,
		getArtistReleases(gresult.Get("releases_url").Str),
		int(gresult.Get("id").Num),
	}

	err = json.NewEncoder(w).Encode(artist)
	if err != nil {
		log.Fatal("jsonEncode:", err)
	}
}

func getArtistReleases(url string) (res []models.ReleaseMin) {
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
}

/**
Search for releases by value
*/
func SearchAlbum(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	method := "album.search"

	url := fmt.Sprintf("%s/?method=%s&album=%s&api_key=%s&format=json", baseUrl, method, params["value"], key)
	//url := fmt.Sprintf("%s/database/search?q=%s&type=master&key=%s&secret=%s", baseUrl, params["value"], key, secret)

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

	var result models.Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println(err)
	}

	var releases []models.Result
	gresult := gjson.Get(bodyString, "results")
	for _, result := range gresult.Array() {

		var styles []string
		for _, res := range result.Get("style").Array() {
			styles = append(styles, res.Str)
		}
		if styles == nil {
			styles = []string{}
		}

		releases = append(releases, models.Result{
			Title:      result.Get("title").Str,
			URI:        result.Get("uri").Str,
			CoverImage: result.Get("thumb").Str,
			ID:         int(result.Get("id").Num),
			Year:       result.Get("year").Str,
			Style:      styles,
		})
	}

	err = json.NewEncoder(w).Encode(releases)
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

	fmt.Printf("get artist: [%s]", params["value"])

	parameter, err := Parse(QueryEscape(params["value"]))
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%s/database/search?q=%s&type=artist&key=%s&secret=%s", baseUrl, parameter.String(), key, secret)

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

	var result models.Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println(err)
	}

	var regex = regexp.MustCompile(`^.*?\([^\d]*(\d+)[^\d]*\).*$`)

	var artists []models.Result
	gresult := gjson.Get(bodyString, "results")
	for _, result := range gresult.Array() {
		if regex.MatchString(result.Get("title").Str) {
			continue
		}

		artists = append(artists, models.Result{
			ID:         int(result.Get("id").Num),
			Title:      result.Get("title").Str,
			URI:        result.Get("uri").Str,
			CoverImage: result.Get("cover_image").Str,
		})
	}

	err = json.NewEncoder(w).Encode(artists)
	if err != nil {
		log.Fatal("jsonEncode:", err)
		return
	}
}
