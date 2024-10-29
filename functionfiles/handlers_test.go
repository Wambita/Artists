package groupie_tracker

import (
	"html/template"
	"io"
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

func TestHomeHandler(t *testing.T) {
	// Create a response recorder to capture the response.
	recorder := httptest.NewRecorder()

	tests := []struct {
		name           string
		method         string
		expectedStatus int
	}{
		{
			name:           "GET request",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "POST request",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "PUT request",
			method:         http.MethodPut,
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new request with the method specified in the test case.
			req, err := http.NewRequest(tt.method, "/", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Reset the response recorder for each test case.
			recorder = httptest.NewRecorder()

			// Call the handler with the recorder and request.
			HomeHandler(recorder, req)

			// Check if the status code matches the expected status.
			if got := recorder.Code; got != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, got)
			}
		})
	}
}

func TestErrorHandler(t *testing.T) {
	recorder := httptest.NewRecorder()

	tests := []struct {
		name           string
		message        string
		statusCode     int
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid error handling",
			message:        "Test error message",
			statusCode:     http.StatusNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Test error message",
		},
		{
			name:           "Internal server error during template execution",
			message:        "Internal Server Error",
			statusCode:     http.StatusInternalServerError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Internal Server Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			recorder = httptest.NewRecorder()
			ErrorHandler(recorder, req, tt.message, tt.statusCode)

			if got := recorder.Code; got != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, got)
			}

			body, err := io.ReadAll(recorder.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			if !strings.Contains(string(body), tt.expectedBody) {
				t.Errorf("Expected body to contain %q, got %q", tt.expectedBody, string(body))
			}
		})
	}
}
