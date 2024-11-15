package groupie_tracker

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	w.WriteHeader(http.StatusOK)

	// pass artists slice to the template
	err := Templates.ExecuteTemplate(w, "index.html", Artists)
	if err != nil {
		ErrorHandler(w, r, "Internal  Server Error", http.StatusInternalServerError)
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
		ErrorHandler(w, r, "Access denied", http.StatusForbidden)
		return
	}

	// Use switch case to handle specific routes

	switch r.URL.Path {
	case "/":
		HomeHandler(w, r)
	case "/artist":
		ArtistHandler(w, r)
	case "/search":
		SearchHandler(w, r)
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
	if len(artistID) > 1 && artistID[0] == '0' {
		ErrorHandler(w, r, "Invalid Artist ID. Leading zeros are not allowed.", http.StatusBadRequest)
		return
	}
	// Parse the artist ID
	id, err := strconv.Atoi(artistID)
	if err != nil || id < 1 || id > len(Artists) {
		ErrorHandler(w, r, "Invalid Artist ID. It must be a number between 0 and 52.", http.StatusBadRequest)
		return
	}

	// Define a struct for the expected data structure
	type datesLocations struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	}

	// Define a struct for the expected data structure
	type concertDates struct {
		ConcertDates []string `json:"dates"`
	}

	// Define a struct for the expected data structure
	type Locations struct {
		Locations []string `json:"locations"`
	}

	var reletions datesLocations

	var locations Locations

	var dates concertDates

	// Artists[id-1]
	if len(Artists[id-1].DatesLocations) == 0 {
		if err := fetchData(RelationURL+artistID, &reletions); err != nil {
			ErrorHandler(w, r, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Artists[id-1].DatesLocations = reletions.DatesLocations
	}

	if len(Artists[id-1].Locations) == 0 {
		if err := fetchData(LocationsURL+artistID, &locations); err != nil {
			ErrorHandler(w, r, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Artists[id-1].Locations = locations.Locations
	}

	if len(Artists[id-1].ConcertDates) == 0 {
		if err := fetchData(DatesURL+artistID, &dates); err != nil {
			ErrorHandler(w, r, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Artists[id-1].ConcertDates = dates.ConcertDates
	}

	// pass the artist into the template
	err1 := Templates.ExecuteTemplate(w, "artist.html", Artists[id-1])

	if err1 != nil {
		ErrorHandler(w, r, "Internal Server Error", http.StatusInternalServerError)
	}
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
		ErrorHandler(w, r, "Internal Server Error", http.StatusInternalServerError)
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
	http.ServeFile(w, r, fullPath)
}

type Suggestion struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if the request path is exactly /search
	if r.URL.Path != "/search" {
		ErrorHandler(w, r, fmt.Sprintf("The Requested path %s does not exist", r.URL.Path), http.StatusNotFound)
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		ErrorHandler(w, r, "search content is required", http.StatusBadRequest)
		return
	}

	query = strings.ToLower(query) // Case-insensitive search
	var suggestions []Suggestion
	// fmt.Println(query)

	for _, artist := range Artists {
		// Search by artist/band name
		if strings.Contains(strings.ToLower(artist.Name), query) {
			suggestions = append(suggestions, Suggestion{ID: artist.ID, Name: artist.Name, Type: "artist/band"})

		}
		// Search by member names
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				suggestions = append(suggestions, Suggestion{ID: artist.ID, Name: member, Type: "member"})

			}
		}
		// Search by locations
		for _, location := range artist.Locations {
			if strings.Contains(strings.ToLower(location), query) {
				suggestions = append(suggestions, Suggestion{ID: artist.ID, Name: location, Type: "location"})

			}
		}
		// Search by creation and first album dates (converting to strings for matching)
		if strings.Contains(fmt.Sprint(artist.Year), query) {
			suggestions = append(suggestions, Suggestion{ID: artist.ID, Name: strconv.Itoa(artist.Year), Type: "creation date"})

		}
		if strings.Contains(fmt.Sprint(artist.Album), query) {
			suggestions = append(suggestions, Suggestion{ID: artist.ID, Name: artist.Album, Type: "first album date"})

		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}
