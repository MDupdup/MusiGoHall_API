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
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do:", err)
		return
	}
	defer resp.Body.Close()

	var artist models.ArtistFull

	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		log.Println(err)
	}

	err = json.NewEncoder(w).Encode(artist)
	if err != nil {
		log.Fatal("jsonEncode:", err)
		return
	}
}

func SearchArtist(w http.ResponseWriter, req *http.Request) {
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
	var artists []models.Artist

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
