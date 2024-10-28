package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	groupie_tracker "groupie-tracker/functionfiles"
)

func main() {
	// Load data
	groupie_tracker.LoadData()

	// Initialize templates
	groupie_tracker.InitializeTemplates()

	// Set routes
	http.HandleFunc("/", groupie_tracker.RouteHandler)
	http.HandleFunc("/artist", groupie_tracker.ArtistHandler)

	// Serve static files
	http.HandleFunc("/static/", groupie_tracker.RouteHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
