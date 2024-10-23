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

// ErrorPage struct
type ErrorPage struct {
	Code    int
	Name    string
	Message string
}