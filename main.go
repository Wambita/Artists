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

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
