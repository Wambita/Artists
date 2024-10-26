package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

	// Serve static files
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		fullPath := filepath.Join("static", r.URL.Path[len("/static/"):])
		// directory restriction
		info, err := os.Stat(fullPath)
		if err == nil && info.IsDir() {
			// http.Error(w, "Forbidden", http.StatusForbidden)
			groupie_tracker.ErrorHandler(w, r, "Forbidden", http.StatusForbidden)
			return
		}
		//
		http.FileServer(http.Dir("static")).ServeHTTP(w, r)
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
