package main

import (
	"LetsGO_API/api"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("magic is happening on port 3333")

	router := mux.NewRouter()

	router.HandleFunc("/search/{type}/{name}", api.SearchArtist).Methods("GET")
	router.HandleFunc("/artist/{id}", api.GetArtistById).Methods("GET")

	log.Fatal(http.ListenAndServe(":3333", router))
}
