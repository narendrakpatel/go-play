package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

// Res struct for response
type Res struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type Queue []string

type Service struct {
	queue Queue
	playChan chan string
	cmd *exec.Cmd
	developerKey string
}

func main() {


	service := Service{}
	service.playChan = make(chan string)

	if err != nil {
		panic(err)
	}
	service.developerKey = string(data)

	router := mux.NewRouter()
	router.HandleFunc("/play", service.QueueSong).Methods("POST")
	router.HandleFunc("/skip", service.SkipSong).Methods("GET")
	router.HandleFunc("/pause", service.PauseSong).Methods("GET")
	router.HandleFunc("/search", service.SearchSong)

	go service.PlaySong()

	log.Fatal(http.ListenAndServe(":1234", router))
}

// QueueSong adds url of song to queue
func (s Service) QueueSong(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")

	s.queue = append(s.queue, url)
	select {
		case s.playChan <- "Hello world":
			fmt.Println("song sent to channel")
		default:
			fmt.Println("song already playing")
	}

}

func (q Queue) append() {

}

func (q Queue) pop() {

}

// SkipSong skips the currently playing song
func (s Service) SkipSong(w http.ResponseWriter, r *http.Request) {
	s.cmd.Process.Kill()
}

// PauseSong pauses the currently playing song
func (Service) PauseSong(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this
}

// ResumeSong resumes the paused song
func (Service) ResumeSong(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this
}

// SearchSong searches songs using YouTube API
func (s Service) SearchSong(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	res := s.getIDs(query)
	enableCors(&w)
	json.NewEncoder(w).Encode(res)
}

// PlaySong plays song in queue
func (s Service) PlaySong() {
	for {
		select {
		case <-s.playChan:
			for len(s.queue) > 0 {
				s.cmd = exec.Command("mpv", "--no-terminal", "--no-video", s.queue[0])
				s.cmd.Run()
				s.queue = s.queue[1:]
			}
		}
	}
}

func initializeService() {
	data, err := ioutil.ReadFile("./api-key")
	maxResults := int64(10)
	client := &http.Client{
		Transport: &transport.APIKey{
			Key: s.developerKey,
		},
	}

	service, err := youtube.New(client)

	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
		return nil
	}

}