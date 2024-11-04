package groupie_tracker

import (
	"fmt"
	"strconv"
	"encoding/json"
	"log"
	"net/http"
)

// Api  urls
// var (
var artistsURL = "https://groupietrackers.herokuapp.com/api/artists"
var LocationsURL = "https://groupietrackers.herokuapp.com/api/locations/"
var DatesURL = "https://groupietrackers.herokuapp.com/api/dates/"
var RelationURL = "https://groupietrackers.herokuapp.com/api/relation/"

// )

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

	for i := range Artists {
		artistID := strconv.Itoa(Artists[i].ID)

		// Define a struct for the expected data structure
		type datesLocations struct {
			DatesLocations map[string][]string `json:"datesLocations"`
		}

		// Define a struct for the expected data structure
		type concertDates struct {
			ConcertDates []string `json:"dates"`
		}

		// Define a struct for the expected data structure
		type Locations struct {
			Locations []string `json:"locations"`
		}

		var reletions datesLocations

		var locations Locations

		var dates concertDates

		// artist
		if len(Artists[i].DatesLocations) == 0{
			if err := fetchData(RelationURL+artistID, &reletions); err != nil {
				// ErrorHandler(w, r, "Internal Server Error", http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
		}		
		Artists[i].DatesLocations = reletions.DatesLocations

	
		if len(Artists[i].Locations) == 0{
			if err := fetchData(LocationsURL+artistID, &locations); err != nil {
				// ErrorHandler(w, r, "Internal Server Error", http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
		}		
		Artists[i].Locations = locations.Locations

		if len(Artists[i].Locations) == 0{
			if err := fetchData(DatesURL+artistID, &dates); err != nil {
				// ErrorHandler(w, r, "Internal Server Error", http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
		}
		
		Artists[i].ConcertDates = dates.ConcertDates
		
	}
}
