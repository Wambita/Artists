package groupie_tracker

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// Api  urls
const (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	datesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	relationURL  = "https://groupietrackers.herokuapp.com/api/relation"
)

// gloabl vars for storing data
var (
	Artists   []Artist
	Locations []Location
	Dates     []Date
	Relations []Relation
	mu        sync.Mutex
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
func loadData() {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		var ar []Artist
		if err := fetchData(artistsURL, &ar); err != nil {
			log.Println("Error fetching artist:", err)
			return
		}
		mu.Lock()
		Artists = ar
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
		Locations = lr.Locations
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
		Dates = dr.Dates
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		var rr RelationResponse
		if err := fetchData(relationURL, &rr); err != nil {
			log.Println("Error fetching relations:", err)
			return
		}
		mu.Lock()
		Relations = rr.Relation
		mu.Unlock()
	}()

	wg.Wait()
}
