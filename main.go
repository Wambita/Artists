package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// "path/filepath"

	groupie_tracker "groupie-tracker/functionfiles"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("usage: go run .")
		os.Exit(0)
	}
	// Load data
	groupie_tracker.LoadData()

	// Initialize templates
	groupie_tracker.InitializeTemplates()

	// Set routes
	http.HandleFunc("/", groupie_tracker.RouteHandler)
	http.HandleFunc("/artist", groupie_tracker.ArtistHandler)

	// Serve static files
	http.HandleFunc("/static/", groupie_tracker.StaticFileHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
