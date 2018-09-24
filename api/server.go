package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

var queue []string
var cmd *exec.Cmd

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/play", QueueSong).Methods("POST")
	router.HandleFunc("/skip", SkipSong).Methods("GET")
	router.HandleFunc("/pause", PauseSong).Methods("GET")

	go PlaySong()

	log.Fatal(http.ListenAndServe(":8000", router))
}

// QueueSong adds url of song to queue
func QueueSong(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if queue == nil {
		fmt.Println("Playing Song...")
	}
	queue = append(queue, url)
}

// PlaySong plays song in queue
func PlaySong() {
	for {
		if queue == nil {
			continue
		}

		cmd = exec.Command("mpv", "--no-terminal", "--no-video", queue[0])
		cmd.Run()

		if len(queue) == 1 {
			queue = nil
		} else {
			queue = queue[1:]
			fmt.Println("Playing next song...")
		}
	}
}

// SkipSong skips the currently playing song
func SkipSong(w http.ResponseWriter, r *http.Request) {
	cmd.Process.Kill()
}

// PauseSong pauses the currently playing song
func PauseSong(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this
}

// ResumeSong resumes the paused song
func ResumeSong(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this
}
