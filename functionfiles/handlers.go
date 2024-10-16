package groupie_tracker

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

// home page handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// pass artists slice to the template
	err := Templates.ExecuteTemplate(w, "index.html", Artists)
	if err != nil {
		http.Error(w, "Internal  Server Error", http.StatusInternalServerError)
	}
}

// artist page handler (indivdual artist)
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artistID := r.URL.Query().Get("id")
	if artistID == "" {
		ErrorHandler(w, r, "Artist ID is required", http.StatusBadRequest)
		return
	}
	for _, artist := range Artists {		
		if fmt.Sprintf("%d", artist.ID) == artistID {
			artist.DatesLocations = reletions(artistID)
			fmt.Println(artist.DatesLocations)
			
			// pass the artist into the template
			err := Templates.ExecuteTemplate(w, "artist.html", artist)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}
	}
}

//function for getting reletions
func reletions(id string) map[string][]string{
	url := "https://groupietrackers.herokuapp.com/api/relation/"+id

	resp, err1 := http.Get(url)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response:", err)
	}

	// Define a struct for the expected data structure
	type ApiResponse struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	}

	// Unmarshal into the struct
	var response ApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("Error unmarshalling:", err)
	}

	// return the DatesLocations map
	return response.DatesLocations
}

// ierror page handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, message string, statusCode int) {

	// error page instance with error details
	errorPage := ErrorPage{
		Code:    statusCode,
		Name:    http.StatusText(statusCode),
		Message: message,
	}

	// set http status code in the response header
	w.WriteHeader(statusCode)

	// render the error page template
	err := Templates.ExecuteTemplate(w, "error.html", errorPage)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
