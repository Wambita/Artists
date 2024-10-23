package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	groupie_tracker "groupie-tracker/functionfiles"
)

func main() {
	// Load data
	groupie_tracker.LoadData()

	// Initialize templates
	groupie_tracker.InitializeTemplates()

	// Set routes
	http.HandleFunc("/", routeHandler)

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

// routeHandler handles requests to the defined routes
func routeHandler(w http.ResponseWriter, r *http.Request) {
	// Restrict to GET method only
	if r.Method != http.MethodGet {
		groupie_tracker.ErrorHandler(w, r, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/static/") {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	// Use switch case to handle specific routes

	switch r.URL.Path {
	case "/":
		groupie_tracker.HomeHandler(w, r)
	case "/artist":
		groupie_tracker.ArtistHandler(w, r)
	default:
		groupie_tracker.ErrorHandler(w, r, fmt.Sprintf("The Requested path %s does not exist", r.URL.Path), http.StatusNotFound)
	}
}
