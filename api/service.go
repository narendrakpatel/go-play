package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

// Service struct provides all services required by the server
type Service struct {
	queue      Queue
	playChan   chan string
	cmd        *exec.Cmd
	maxResults int64
	client     *youtube.Service
}

// initService initializes the basic requirements of the server
func initService() Service {
	data, err := ioutil.ReadFile("./api-key")
	if err != nil {
		panic(err)
	}

	developerKey := string(data)
	maxResults := int64(10)
	client := &http.Client{
		Transport: &transport.APIKey{
			Key: developerKey,
		},
	}

	youtubeService, err := youtube.New(client)

	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
		return Service{}
	}

	return Service{
		queue:      Queue{},
		playChan:   make(chan string),
		cmd:        nil,
		maxResults: maxResults,
		client:     youtubeService,
	}

}

// QueueSong adds url of song to queue
func (s *Service) QueueSong(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	url := r.URL.Query().Get("url")
	s.queue.push(url)

	select {
	case s.playChan <- "Play":
		fmt.Println("Song send for playing")
	default:
		fmt.Println("Sent to queue")
	}
}

// SkipSong skips the currently playing song
func (s *Service) SkipSong(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	s.cmd.Process.Kill()
}

// PauseSong pauses the currently playing song
func (s Service) PauseSong(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this
}

// ResumeSong resumes the paused song
func (s Service) ResumeSong(w http.ResponseWriter, r *http.Request) {
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
func (s *Service) PlaySong() {
	for {
		select {
		case <-s.playChan:
			for len(s.queue) > 0 {
				s.cmd = exec.Command("mpv", "--no-terminal", "--no-video", s.queue[0])
				s.cmd.Run()
				s.queue.pop()
			}
		}
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
