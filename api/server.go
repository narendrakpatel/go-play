package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	service := initService()

	router := mux.NewRouter()
	router.HandleFunc("/play", service.QueueSong).Methods("POST")
	router.HandleFunc("/skip", service.SkipSong).Methods("GET")
	router.HandleFunc("/pause", service.PauseSong).Methods("GET")
	router.HandleFunc("/search", service.SearchSong)

	go service.PlaySong()

	log.Fatal(http.ListenAndServe(":8000", router))
}
