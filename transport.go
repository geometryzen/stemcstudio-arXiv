package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type searchRequest struct {
	Query string `json:"query"`
}

type searchRef struct {
	HRef     string   `json:"href"`
	Owner    string   `json:"owner"`
	GistID   string   `json:"gistId"`
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	Keywords []string `json:"keywords"`
}

type searchResponse struct {
	Found int64       `json:"found"`
	Start int64       `json:"start"`
	Refs  []searchRef `json:"refs,omitempty"`
}

type submissionPayload struct {
	Owner    string   `json:"owner"`
	GistID   string   `json:"gistId"`
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	Keywords []string `json:"keywords"`
}

func makeSearchHandlerFunc(service SearchService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload searchRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&payload)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		fmt.Printf("Searching for %s\n", payload.Query)
		finds, err := service.Search(payload.Query, 30)
		if err != nil {
			fmt.Println("err  : ", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		response := searchResponse{Found: finds.Found, Start: finds.Start}
		// It's a bit wierd that we have to copy the data to get the serialization to come out right.
		// But maybe we should accept that our service and serialized format are loosely coupled.
		for _, record := range finds.Refs {
			href := record.HRef
			owner := record.Owner
			gistID := record.GistID
			title := record.Title
			author := record.Author
			keywords := record.Keywords
			response.Refs = append(response.Refs, searchRef{Author: author, GistID: gistID, HRef: href, Owner: owner, Title: title, Keywords: keywords})
		}
		json.NewEncoder(w).Encode(&response)
	}
}

func makeSubmitHandlerFunc(service SearchService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload submissionPayload
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&payload)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		packet := Submission{Author: payload.Author, GistID: payload.GistID, Keywords: payload.Keywords, Owner: payload.Owner, Title: payload.Title}
		_, err = service.Submit(&packet)
		if err != nil {
			fmt.Println("err  : ", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		http.Error(w, "OK", http.StatusOK)
	}
}
