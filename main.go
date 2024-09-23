package main

import (
    "encoding/json"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"
    "sync"
)

// API URLs
const (
    artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
    locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
    datesURL     = "https://groupietrackers.herokuapp.com/api/dates"
    relationURL  = "https://groupietrackers.herokuapp.com/api/relation"
)

// Structs matching the JSON structure
type Artist struct {
    ID      int      `json:"id"`
    Name    string   `json:"name"`
    Image   string   `json:"image"`
    Year    int      `json:"year"`
    Album   string   `json:"first_album"`
    Members []string `json:"members"`
}

type Location struct {
    ID   int    `json:"id"`
    City string `json:"city"`
}

type Date struct {
    ID   int    `json:"id"`
    Date string `json:"date"`
}

type Relation struct {
    ArtistID   int `json:"artist_id"`
    LocationID int `json:"location_id"`
    DateID     int `json:"date_id"`
}

// Wrapper structs for API responses
type ArtistsResponse struct {
    Artists []Artist `json:"artists"`
}

type LocationsResponse struct {
    Locations []Location `json:"locations"`
}

type DatesResponse struct {
    Dates []Date `json:"dates"`
}

type RelationResponse struct {
    Relation []Relation `json:"relation"`
}

// Global variables
var (
    artists   []Artist
    locations []Location
    dates     []Date
    relation  []Relation
    mu        sync.Mutex
    templates *template.Template
)

// FetchData fetches JSON data from a URL and decodes it into the target interface
func fetchData(url string, target interface{}) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    return json.NewDecoder(resp.Body).Decode(target)
}
// LoadData concurrently fetches data from all APIs
func loadData() {
    var wg sync.WaitGroup
    wg.Add(4)

    go func() {
        defer wg.Done()
        var ar []Artist  // Change to a slice of Artist
        if err := fetchData(artistsURL, &ar); err != nil {
            log.Println("Error fetching artists:", err)
            return
        }
        mu.Lock()
        artists = ar  // Directly assign the slice to artists
        mu.Unlock()
    }()

    go func() {
        defer wg.Done()
        var lr LocationsResponse
        if err := fetchData(locationsURL, &lr); err != nil {
            log.Println("Error fetching locations:", err)
            return
        }
        mu.Lock()
        locations = lr.Locations
        mu.Unlock()
    }()

    go func() {
        defer wg.Done()
        var dr DatesResponse
        if err := fetchData(datesURL, &dr); err != nil {
            log.Println("Error fetching dates:", err)
            return
        }
        mu.Lock()
        dates = dr.Dates
        mu.Unlock()
    }()

    go func() {
        defer wg.Done()
        var rr RelationResponse
        if err := fetchData(relationURL, &rr); err != nil {
            log.Println("Error fetching relation:", err)
            return
        }
        mu.Lock()
        relation = rr.Relation
        mu.Unlock()
    }()

    wg.Wait()
}


// InitializeTemplates parses all template files at startup
func InitializeTemplates() {
    var err error
    templates, err = template.ParseGlob("templates/*.html")
    if err != nil {
        log.Fatalf("Error parsing templates: %v", err)
    }
}

// Handler for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    // Pass the artists slice to the template
    err := templates.ExecuteTemplate(w, "index.html", artists)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Handler for individual artist pages
func artistHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    artistID := r.URL.Query().Get("id")
    for _, artist := range artists {
        if fmt.Sprintf("%d", artist.ID) == artistID {
            // Pass the single artist to the template
            err := templates.ExecuteTemplate(w, "artist.html", artist)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
            return
        }
    }
    http.Error(w, "Artist not found", http.StatusNotFound)
}

func main() {
    // Load data
    loadData()

    // Initialize templates
    InitializeTemplates()

    // Set up routes
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/artist", artistHandler)

    // Serve static files (e.g., CSS, images) from the "static" directory
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Start the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    fmt.Printf("Server starting on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
