package main

import (
	"MusiGoHall_API/api"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("magic is happening on port 3333")

	router := mux.NewRouter()

	router.HandleFunc("/search/release/{value}", api.SearchRelease).Methods("GET")
	router.HandleFunc("/search/artist/{value}", api.SearchArtist).Methods("GET")
	router.HandleFunc("/search/label/{value}", api.SearchLabel).Methods("GET")

	router.HandleFunc("/release/{id}", api.GetRelease).Methods("GET")
	router.HandleFunc("/artist/{id}", api.GetArtist).Methods("GET")
	router.HandleFunc("/label/{id}", api.GetLabel).Methods("GET")

	log.Fatal(http.ListenAndServe(":3333", router))
}
