package api

import (
	"LetsGO_API/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var baseUrl = "https://api.discogs.com"
var key = "hkPZgJqVNACgrhiCUBEr"
var secret = "GhhVfKpuGbRJhcktPLNctfTXpzKUvpKl"

func GetArtistById(w http.ResponseWriter, req *http.Request) {
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

func ComplexSearch(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	url := buildURL(params)

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

	switch params["type"] {
	case "release":
		var model []models.Release
		var result models.Result

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Println(err)
		}

		for i := 0; i < len(result.Results); i++ {
			//newModel := result.Results[i]
			//model = append(model, newModel)
		}

		err = json.NewEncoder(w).Encode(model)
		if err != nil {
			log.Fatal("jsonEncode:", err)
			return
		}
		break
	case "artist":
		var model []models.Artist
		var result models.Result

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Println(err)
		}

		for i := 0; i < len(result.Results); i++ {
			newModel := result.Results[i]
			model = append(model, newModel)
		}

		err = json.NewEncoder(w).Encode(model)
		if err != nil {
			log.Fatal("jsonEncode:", err)
			return
		}
		break
	case "label":
		var model []models.Label
		var result models.Result

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Println(err)
		}

		for i := 0; i < len(result.Results); i++ {
			newModel := result.Results[i]
			model = append(model, newModel)
		}

		err = json.NewEncoder(w).Encode(model)
		if err != nil {
			log.Fatal("jsonEncode:", err)
			return
		}
		break
	default:
		break
	}

}

func SearchRelease(w http.ResponseWriter, req *http.Request) {
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
	var releases []models.Release

	//decoder := json.NewDecoder(resp.Body)

	//var lol = []byte

	//println(resp.Body.Read(lol))

	for i := 0; i < len(result.Results); i++ {
		//releases = append(releases, result.Results[i])
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
