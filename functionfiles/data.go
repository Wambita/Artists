package groupie_tracker

import (
	"encoding/json"
	"log"
	"net/http"
)

// Api  urls
var (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	datesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	relationURL  = "https://groupietrackers.herokuapp.com/api/relation"
)

// gloabl vars for storing data
var (
	Artists []Artist
)

// fetch json data form a url and decodes ot into target interface
func fetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

// loadData fetch data concurrently from all apis
func LoadData() {
	var ar []Artist
	if err := fetchData(artistsURL, &ar); err != nil {
		log.Println("Error fetching artist:", err)
		return
	}
	Artists = ar
}

// Function to fetch locations from the URL
func fetchLocations(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Define a struct for the expected response
	type LocationsResponse struct {
		Locations []string `json:"locations"` // Adjust the JSON key based on your API response structure
	}

	var locationsResponse LocationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationsResponse); err != nil {
		return nil, err
	}

	return locationsResponse.Locations, nil
}

// Function to fetch concert dates from the URL
func fetchDates(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Define a struct for the expected response
	var dates []string
	if err := json.NewDecoder(resp.Body).Decode(&dates); err != nil {
		return nil, err
	}

	return dates, nil
}
