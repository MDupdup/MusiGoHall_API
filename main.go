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
	router.HandleFunc("/album/{id}/{name}", api.GetAlbum).Methods("GET")
	router.HandleFunc("/artist/{id}", api.GetArtist).Methods("GET")

	router.HandleFunc("/favorites/add/{type}", api.AddToDB).Methods("POST")
	router.HandleFunc("/favorites/get/{type}", api.GetDB).Methods("GET")
	router.HandleFunc("/favorites/del/{type}/{value}", api.DeleteFromDB).Methods("GET")

	router.HandleFunc("/favorites/isindb/{type}/{name}", api.IsInDB).Methods("GET")

	log.Fatal(http.ListenAndServe(":3333", router))
}
