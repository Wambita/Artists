package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "strings"

	groupie_tracker "groupie-tracker/functionfiles"
)

func main() {
	// load data
	groupie_tracker.LoadData()

	// initialize templates
	groupie_tracker.InitializeTemplates()

	// set routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		switch path{
		case "/":
			groupie_tracker.HomeHandler(w, r)
		case "/artist":
			if r.URL.Query().Get("id") == "" {
				groupie_tracker.ErrorHandler(w, r, "Artist ID is required", http.StatusBadRequest)
				return
			}
			groupie_tracker.ArtistHandler(w, r)
		default:
			groupie_tracker.ErrorHandler(w, r, fmt.Sprintf("The Requested path  %s  does not exist", path), http.StatusNotFound)
		}

		// if path == "/" {
		// 	groupie_tracker.HomeHandler(w, r)
		// 	return
		// }

		// if strings.HasPrefix(path, "/artist") {
		// 	if r.URL.Query().Get("id") == "" {
		// 		groupie_tracker.ErrorHandler(w, r, "Artist ID is required", http.StatusBadRequest)
		// 		return
		// 	}
		// 	groupie_tracker.ArtistHandler(w, r)
		// 	return
		// }
		// // invalid routes
		// groupie_tracker.ErrorHandler(w, r, fmt.Sprintf("The Requested path  %s  does not exist", path), http.StatusNotFound)
	})

	// server static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
