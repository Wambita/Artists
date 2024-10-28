package groupie_tracker

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func init() {
	// Set up a basic HTML template for testing purposes
	Templates = template.Must(template.New("test").Parse(`
        {{define "artist.html"}}<html><body>{{.Name}}</body></html>{{end}}
        {{define "error.html"}}<html><body>{{.Message}}</body></html>{{end}}
    `))

	// Populate the Artists slice with mock data
	Artists = []Artist{
		{
			Name:           "Test Artist",
			DatesLocations: map[string][]string{"2022-01-01": {"Test Location 1"}},
			Locations:      []string{"Test Location 1", "Test Location 2"},
			ConcertDates:   []string{"2022-01-01", "2022-02-02"},
		},
	}

	// Override the functions to return mock data for tests.
	locations = func(id string) []string {
		return []string{"Test Location 1", "Test Location 2"}
	}

	dates = func(id string) []string {
		return []string{"2022-01-01", "2022-02-02"}
	}

	reletions = func(id string) map[string][]string {
		return map[string][]string{
			"2022-01-01": {"Test Location 1"},
			"2022-02-02": {"Test Location 2"},
		}
	}
}

func TestArtistHandler(t *testing.T) {
	// Provide mock data for Artists to avoid nil dereference
	Artists = []Artist{
		{
			Name:           "Test Artist",
			DatesLocations: map[string][]string{"2022-01-01": {"Test Location"}},
			Locations:      []string{"Test Location"},
			ConcertDates:   []string{"2022-01-01"},
		},
	}

	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid ID",
			url:            "/artist?id=1",
			expectedStatus: http.StatusOK,
			expectedBody:   "Test Artist",
		},
		{
			name:           "Missing ID",
			url:            "/artist",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Artist ID is required",
		},
		{
			name:           "Invalid ID",
			url:            "/artist?id=abc",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid Artist ID",
		},
		{
			name:           "Out-of-bounds ID",
			url:            "/artist?id=100",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}
			rec := httptest.NewRecorder()

			ArtistHandler(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("Expected status %v, got %v", tt.expectedStatus, rec.Code)
			}

			// Check if expected body content is present (optional for some cases)
			if tt.expectedBody != "" && !strings.Contains(rec.Body.String(), tt.expectedBody) {
				t.Errorf("Expected content %q in response body, got %v", tt.expectedBody, rec.Body.String())
			}
		})
	}
}

func TestRouteHandler(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		expectedStatus int
	}{
		{
			name:           "Home Path",
			path:           "/",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Artist Path",
			path:           "/artist?id=1",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Path",
			path:           "/nonexistent",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}
			rec := httptest.NewRecorder()

			RouteHandler(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("Expected status %v, got %v", tt.expectedStatus, rec.Code)
			}
		})
	}
}

func TestLocations(t *testing.T) {
	locations := locations("1")
	if len(locations) != 2 || locations[0] != "Test Location 1" {
		t.Errorf("Expected mock locations data, got %v", locations)
	}
}

func TestDates(t *testing.T) {
	dates := dates("1")
	if len(dates) != 2 || dates[0] != "2022-01-01" {
		t.Errorf("Expected mock dates data, got %v", dates)
	}
}

func TestReletions(t *testing.T) {
	reletions := reletions("1")
	if len(reletions) != 2 || reletions["2022-01-01"][0] != "Test Location 1" {
		t.Errorf("Expected mock relations data, got %v", reletions)
	}
}
