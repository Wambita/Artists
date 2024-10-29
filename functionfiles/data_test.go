package groupie_tracker

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Artist1 struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

// Mock response for artists
var mockArtistsResponse = []Artist1{
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

func TestFetchData(t *testing.T) {
	// Set up a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/artists":
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mockArtistsResponse)
		case "/locations":
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mockLocationsResponse)
		case "/dates":
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mockDatesResponse)
		case "/notfound":
			w.WriteHeader(http.StatusNotFound)
		case "/invalidjson":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("invalid json"))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	tests := []struct {
		name    string
		url     string
		target  interface{}
		wantErr bool
	}{
		{
			name:    "Fetch Artists",
			url:     ts.URL + "/artists",
			target:  &[]Artist1{},
			wantErr: false,
		},
		{
			name:    "Fetch Locations",
			url:     ts.URL + "/locations",
			target:  &struct{ Locations []string }{},
			wantErr: false,
		},
		{
			name:    "Fetch Dates",
			url:     ts.URL + "/dates",
			target:  &[]string{},
			wantErr: false,
		},
		{
			name:    "404 Not Found",
			url:     ts.URL + "/notfound",
			target:  &[]Artist1{},
			wantErr: true,
		},
		{
			name:    "Invalid JSON response",
			url:     ts.URL + "/invalidjson",
			target:  &[]Artist1{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fetchData(tt.url, tt.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchData() error = %v, wantErr %v", err, tt.wantErr)
			}

			// If no error is expected, verify the data
			if !tt.wantErr {
				switch tt.url {
				case ts.URL + "/artists":
					var artists []Artist1
					if err := fetchData(tt.url, &artists); err == nil {
						if len(artists) != len(mockArtistsResponse) {
							t.Errorf("Expected %d artists, got %d", len(mockArtistsResponse), len(artists))
						}

						for i, artist := range artists {
							if artist != mockArtistsResponse[i] {
								t.Errorf("Expected artist %v, got %v", mockArtistsResponse[i], artists[i])
							}
						}
					}
				case ts.URL + "/locations":
					var locations struct{ Locations []string }
					if err := fetchData(tt.url, &locations); err == nil {
						if len(locations.Locations) != len(mockLocationsResponse.Locations) {
							t.Errorf("Expected %d locations, got %d", len(mockLocationsResponse.Locations), len(locations.Locations))
						}
					}
				case ts.URL + "/dates":
					var dates []string
					if err := fetchData(tt.url, &dates); err == nil {
						if len(dates) != len(mockDatesResponse) {
							t.Errorf("Expected %d dates, got %d", len(mockDatesResponse), len(dates))
						}
					}
				}
			}
		})
	}
}
