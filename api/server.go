package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	service := initService()

	router := mux.NewRouter()
	router.HandleFunc("/play", service.QueueSong).Methods("POST")
	router.HandleFunc("/skip", service.SkipSong).Methods("GET")
	router.HandleFunc("/pause", service.PauseSong).Methods("GET")
	router.HandleFunc("/search", service.SearchSong)

	go service.PlaySong()

	log.Fatal(http.ListenAndServe(":1234", router))
}
