package main

import (
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

func (s Service) getIDs(q string) []Res {
	call := service.Search.List("id,snippet").
		Q(q).
		MaxResults(maxResults)
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

	return s.parseIDs(videos)
}

func (s Service) parseIDs(matches map[string]string) []Res {
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
