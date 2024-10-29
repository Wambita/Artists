package groupie_tracker

import (
	"encoding/json"
	"log"
	"net/http"
)

// Api  urls
var (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsURL = "https://groupietrackers.herokuapp.com/api/locations/"
	DatesURL     = "https://groupietrackers.herokuapp.com/api/dates/"
	RelationURL  = "https://groupietrackers.herokuapp.com/api/relation/"
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