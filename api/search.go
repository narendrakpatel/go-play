package main

import (
	"log"
	"net/http"
)

// Res struct for response
type Res struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func (s Service) getIDs(q string) []Res {
	call := s.client.Search.List("id,snippet").
		Q(q).
		MaxResults(s.maxResults)
	response, err := call.Do()

	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
		return nil
	}

	videos := make(map[string]string)

	for _, item := range response.Items {
		if item.Id.Kind == "youtube#video" {
			videos[item.Id.VideoId] = item.Snippet.Title
		}
	}

	return parseIDs(videos)
}

func parseIDs(matches map[string]string) []Res {
	var res []Res
	for id := range matches {
		res = append(res, Res{ID: id, URL: getURL(id)})
	}
	return res
}

func getURL(id string) string {
	return "https://www.youtube.com/watch?v=" + id
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
