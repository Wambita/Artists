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

	req, err := http.NewRequest("GET", "/artist?id=1", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rec := httptest.NewRecorder()

	ArtistHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %v", rec.Code)
	}

	// Optional: Check if the response body contains expected data
	expectedContent := "Test Artist"
	if !strings.Contains(rec.Body.String(), expectedContent) {
		t.Errorf("Expected content %q in response body, got %v", expectedContent, rec.Body.String())
	}
}

func TestArtistHandler_MissingID(t *testing.T) {
	req, err := http.NewRequest("GET", "/artist", nil) // Missing ID
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rec := httptest.NewRecorder()

	ArtistHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %v", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Artist ID is required") {
		t.Errorf("Expected error message, got %v", rec.Body.String())
	}
}

func TestArtistHandler_InvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/artist?id=abc", nil) // Non-numeric ID
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rec := httptest.NewRecorder()

	ArtistHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %v", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Invalid Artist ID") {
		t.Errorf("Expected invalid ID error message, got %v", rec.Body.String())
	}
}

func TestArtistHandler_OutOfBoundsID(t *testing.T) {
	req, err := http.NewRequest("GET", "/artist?id=100", nil) // Out-of-bounds ID
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rec := httptest.NewRecorder()

	ArtistHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %v", rec.Code)
	}
}

// func TestRouteHandler_HomePath(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/", nil)
// 	if err != nil {
// 		t.Fatalf("Could not create request: %v", err)
// 	}
// 	rec := httptest.NewRecorder()

// 	RouteHandler(rec, req)

// 	if rec.Code != http.StatusOK {
// 		t.Errorf("Expected status 200, got %v", rec.Code)
// 	}
// }

func TestRouteHandler_ArtistPath(t *testing.T) {
	req, err := http.NewRequest("GET", "/artist?id=1", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rec := httptest.NewRecorder()

	RouteHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %v", rec.Code)
	}
}

func TestRouteHandler_InvalidPath(t *testing.T) {
	req, err := http.NewRequest("GET", "/nonexistent", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rec := httptest.NewRecorder()

	RouteHandler(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %v", rec.Code)
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
