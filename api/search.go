package main

import (
	"google.golang.org/api/youtube/v3"
	"log"
)

// Res struct for response
type Res struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	Thumbnail string `json:"thumb"`
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

	videos := make(map[string]*youtube.SearchResultSnippet)

	for _, item := range response.Items {
		if item.Id.Kind == "youtube#video" {
			videos[item.Id.VideoId] = item.Snippet
		}
	}

	return parseIDs(videos)
}

func parseIDs(matches map[string]*youtube.SearchResultSnippet) []Res {
	var res []Res
	for id, snippet := range matches {
		res = append(res, Res{Title: snippet.Title,
			URL: "https://www.youtube.com/watch?v=" + id,
			Thumbnail: snippet.Thumbnails.High.Url},
			)
	}
	return res
}
