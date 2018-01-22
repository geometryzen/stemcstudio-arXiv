package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	flag.Parse()

	router := mux.NewRouter()

	// The following lines
	fmt.Printf("AWS_ACCESS_KEY_ID => %s\n", os.Getenv("AWS_ACCESS_KEY_ID"))
	fmt.Printf("len(AWS_SECRET_ACCESS_KEY) => %d\n", len(os.Getenv("AWS_SECRET_ACCESS_KEY")))

	searchService := NewSearchService()

	router.HandleFunc("/search", makeSearchHandlerFunc(searchService))
	router.HandleFunc("/submissions", makeSubmitHandlerFunc(searchService))

	server := &http.Server{
		Addr:    "0.0.0.0:8081",
		Handler: router,
	}

	server.ListenAndServe()
}
