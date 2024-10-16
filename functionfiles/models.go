package groupie_tracker

// Artist struct
type Artist struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Year    int      `json:"creationDate"`
	Album   string   `json:"firstAlbum"`
	Members []string `json:"members"`
	DatesLocations map[string][]string
}

// location struct
type Location struct {
	ID   int    `json:"id"`
	City string `json:"city"`
}

// concert date struct
type Date struct {
	ID   int    `json:"id"`
	Date string `json:"date"`
}

// relation struct btwn artist, location and date
type Relation struct {
	ArtistID   int `json:"artist_id"`
	LocationID int `json:"location_id"`
	DateID     int `json:"date_id"`
}

// ErrorPage struct
type ErrorPage struct {
	Code    int
	Name    string
	Message string
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
