package groupie_tracker

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock response for artists
var mockArtistsResponse = []Artist{
	{ID: 1, Name: "Artist One", Image: "image1.jpg"},
	{ID: 2, Name: "Artist Two", Image: "image2.jpg"},
}

// Mock response for locations
var mockLocationsResponse = struct {
	Locations []string `json:"locations"`
}{
	Locations: []string{"Location A", "Location B"},
}

// Mock response for dates
var mockDatesResponse = []string{"2024-10-30", "2024-11-05"}

// // Test fetchData
// func TestFetchData(t *testing.T) {
// 	// Setup a mock server
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.URL.String() == artistsURL {
// 			json.NewEncoder(w).Encode(mockArtistsResponse)
// 		} else {
// 			http.NotFound(w, r)
// 		}
// 	}))
// 	defer server.Close()

// 	// Update the URL to the mock server
// 	artistsURL = server.URL

// 	var artists []Artist
// 	err := fetchData(artistsURL, &artists)
// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}

// 	if len(artists) != len(mockArtistsResponse) {
// 		t.Fatalf("Expected %d artists, got %d", len(mockArtistsResponse), len(artists))
// 	}
// }

// Test LoadData
func TestLoadData(t *testing.T) {
	// Setup a mock server for artists
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(mockArtistsResponse)
	}))
	defer server.Close()

	// Update the URL to the mock server
	artistsURL = server.URL

	LoadData()

	if len(Artists) != len(mockArtistsResponse) {
		t.Fatalf("Expected %d artists, got %d", len(mockArtistsResponse), len(Artists))
	}
}

// Test fetchLocations
// func TestFetchLocations(t *testing.T) {
// 	// Setup a mock server for locations
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.URL.String() == locationsURL {
// 			json.NewEncoder(w).Encode(mockLocationsResponse)
// 		} else {
// 			http.NotFound(w, r)
// 		}
// 	}))
// 	defer server.Close()

// 	// Update the URL to the mock server
// 	locationsURL = server.URL

// 	locations, err := fetchLocations(locationsURL)
// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}

// 	if len(locations) != len(mockLocationsResponse.Locations) {
// 		t.Fatalf("Expected %d locations, got %d", len(mockLocationsResponse.Locations), len(locations))
// 	}
// }

// Test fetchDates
// func TestFetchDates(t *testing.T) {
// 	// Setup a mock server for dates
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.URL.String() == datesURL {
// 			json.NewEncoder(w).Encode(mockDatesResponse)
// 		} else {
// 			http.NotFound(w, r)
// 		}
// 	}))
// 	defer server.Close()

// 	// Update the URL to the mock server
// 	datesURL = server.URL

// 	dates, err := fetchDates(datesURL)
// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}

// 	if len(dates) != len(mockDatesResponse) {
// 		t.Fatalf("Expected %d dates, got %d", len(mockDatesResponse), len(dates))
// 	}
// }
