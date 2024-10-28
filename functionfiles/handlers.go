package groupie_tracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// home page handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// pass artists slice to the template
	err := Templates.ExecuteTemplate(w, "index.html", Artists)
	if err != nil {
		http.Error(w, "Internal  Server Error", http.StatusInternalServerError)
	}
}

// routeHandler handles requests to the defined routes
func RouteHandler(w http.ResponseWriter, r *http.Request) {
	// Restrict to GET method only
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/static/") {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	// Use switch case to handle specific routes

	switch r.URL.Path {
	case "/":
		HomeHandler(w, r)
	case "/artist":
		ArtistHandler(w, r)
	default:
		ErrorHandler(w, r, fmt.Sprintf("The Requested path %s does not exist", r.URL.Path), http.StatusNotFound)
	}
}

// artist page handler (indivdual artist)
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if the request path is exactly /artist
	if r.URL.Path != "/artist" {
		ErrorHandler(w, r, fmt.Sprintf("The Requested path %s does not exist", r.URL.Path), http.StatusNotFound)
		return
	}

	artistID := r.URL.Query().Get("id")
	if artistID == "" {
		ErrorHandler(w, r, "Artist ID is required", http.StatusBadRequest)
		return
	}
	// Parse the artist ID
	id, err := strconv.Atoi(artistID)
	if err != nil || id < 1 || id > 52 {
		ErrorHandler(w, r, "Invalid Artist ID. It must be a number between 0 and 52.", http.StatusBadRequest)
		return
	}

	// Artists[id-1]
	if len(Artists[id-1].DatesLocations) == 0 {
		Artists[id-1].DatesLocations = reletions(artistID)
	}

	if len(Artists[id-1].Locations) == 0 {
		Artists[id-1].Locations = locations(artistID)
	}

	if len(Artists[id-1].ConcertDates) == 0 {
		Artists[id-1].ConcertDates = dates(artistID)
	}

	// pass the artist into the template
	err1 := Templates.ExecuteTemplate(w, "artist.html", Artists[id-1])

	if err1 != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// function for getting locations
func locations(id string) []string {
	url := "https://groupietrackers.herokuapp.com/api/locations/" + id

	resp, err1 := http.Get(url)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}

	// Define a struct for the expected data structure
	type ApiResponse struct {
		Locations []string `json:"locations"`
	}

	// Unmarshal into the struct
	var response ApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("Error unmarshalling:", err)
	}

	// return the DatesLocations map
	return response.Locations
}

// function for getting  dates
func dates(id string) []string {
	url := "https://groupietrackers.herokuapp.com/api/dates/" + id

	resp, err1 := http.Get(url)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}

	// Define a struct for the expected data structure
	type ApiResponse struct {
		ConcertDates []string `json:"dates"`
	}

	// Unmarshal into the struct
	var response ApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("Error unmarshalling:", err)
	}

	// return the DatesLocations map
	return response.ConcertDates
}

// function for getting reletions
func reletions(id string) map[string][]string {
	url := "https://groupietrackers.herokuapp.com/api/relation/" + id

	resp, err1 := http.Get(url)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}

	// Define a struct for the expected data structure
	type ApiResponse struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	}

	// Unmarshal into the struct
	var response ApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("Error unmarshalling:", err)
	}

	// return the DatesLocations map
	return response.DatesLocations
}

// ierror page handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, message string, statusCode int) {
	// error page instance with error details
	errorPage := ErrorPage{
		Code:    statusCode,
		Name:    http.StatusText(statusCode),
		Message: message,
	}

	// set http status code in the response header
	w.WriteHeader(statusCode)

	// render the error page template
	err := Templates.ExecuteTemplate(w, "error.html", errorPage)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func StaticFileHandler(w http.ResponseWriter, r *http.Request) {
	// Prevent direct access to the static directory
	if r.URL.Path == "/static/" || r.URL.Path[len("/static/"):] == "" {
		ErrorHandler(w, r, "Access denied", http.StatusForbidden)
		return
	}

	// Serve the requested static file
	fullPath := filepath.Join("static", r.URL.Path[len("/static/"):])

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		fmt.Println("error: static file not found")
		return
	}

	http.ServeFile(w, r, fullPath)
}
