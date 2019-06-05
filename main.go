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

	router.HandleFunc("/search/album/{value}", api.SearchAlbum).Methods("GET")
	router.HandleFunc("/search/artist/{value}", api.SearchArtist).Methods("GET")

	router.HandleFunc("/album/{id}", api.GetAlbum).Methods("GET")
	router.HandleFunc("/artist/{id}", api.GetArtist).Methods("GET")

	router.HandleFunc("/favorites/add", api.AddToDB).Methods("POST")
	router.HandleFunc("/favorites/all", api.GetDB).Methods("GET")

	log.Fatal(http.ListenAndServe(":3333", router))
}
