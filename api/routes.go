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
)

var baseUrl = "https://api.discogs.com"
var key = "hkPZgJqVNACgrhiCUBEr"
var secret = "GhhVfKpuGbRJhcktPLNctfTXpzKUvpKl"

func GetRelease(w http.ResponseWriter, req *http.Request) {
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
		log.Fatal("Do:", err)
	}
	defer resp.Body.Close()

	var artist models.Release

	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		log.Println(err)
	}

	err = json.NewEncoder(w).Encode(artist)
	if err != nil {
		log.Fatal("jsonEncode:", err)
	}
}

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
		log.Fatal("Do:", err)
	}
	defer resp.Body.Close()

	var artist models.ArtistFull

	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		log.Println(err)
	}

	err = json.NewEncoder(w).Encode(artist)
	if err != nil {
		log.Fatal("jsonEncode:", err)
	}
}

func GetLabel(w http.ResponseWriter, req *http.Request) {
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
		log.Fatal("Do:", err)
	}
	defer resp.Body.Close()

	var artist models.ArtistFull

	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		log.Println(err)
	}

	err = json.NewEncoder(w).Encode(artist)
	if err != nil {
		log.Fatal("jsonEncode:", err)
	}
}

func SearchRelease(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	url := fmt.Sprintf("%s/database/search?q=%s&type=release&key=%s&secret=%s", baseUrl, params["name"], key, secret)

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

	var releases []models.Release
	gresult := gjson.Get(bodyString, "results")
	for _, result := range gresult.Array() {

		println(result.Array())

		releases = append(releases, models.Release{
			Style: []result.,
			MasterID: int(result.Get("master_id").Num),
			Thumb: result.Get("thumb").Str,
			Format: nil,
			Country: result.Get("country").Str,
			Barcode: nil,
			URI: result.Get("uri").Str,
			MasterURL: result.Get("master_url").Str,
				})
	}

	err = json.NewEncoder(w).Encode(releases)
	if err != nil {
		log.Fatal("jsonEncode:", err)
		return
	}
}

func SearchLabel(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	fmt.Println(params["type"], ":", params["name"])
	url := fmt.Sprintf("%s/database/search?q=%s&type=%s&key=%s&secret=%s", baseUrl, params["name"], params["type"], key, secret)

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

	var result models.Result
	var artists []models.Label

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println(err)
	}

	for i := 0; i < len(result.Results); i++ {
		newArtist := result.Results[i]
		artists = append(artists, newArtist)
	}

	err = json.NewEncoder(w).Encode(artists)
	if err != nil {
		log.Fatal("jsonEncode:", err)
		return
	}
}

func buildURL(requestParams map[string]string) (val string) {
	url := fmt.Sprintf("%s/database/search?", baseUrl)

	params := [18]string{
		"query",
		"type",
		"title",
		"release_title",
		"credit",
		"artist",
		"anv",
		"label",
		"genre",
		"style",
		"country",
		"year",
		"format",
		"catno",
		"barcode",
		"track",
		"submitter",
		"contributor",
	}

	for i := 0; i < len(params); i++ {
		if params[i] == "query" {
			url += "q=" + requestParams[params[i]]
			continue
		}
		if requestParams[params[i]] != "" {
			url += "&" + params[i] + "=" + requestParams[params[i]]
		}
	}
	url += "&key=" + key + "&secret=" + secret

	return url
}
