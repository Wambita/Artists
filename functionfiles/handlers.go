package groupie_tracker

import (
	"fmt"
	"net/http"
)

// home page handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// pass artists slice to the template
	err := Templates.ExecuteTemplate(w, "index.html", Artists)
	if err != nil {
		http.Error(w, "Internal  Server Error", http.StatusInternalServerError)
	}
}

// artist page handler (indivdual artist)
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	artistID := r.URL.Query().Get("id")
	for _, artist := range Artists {
		if fmt.Sprintf("%d", artist.ID) == artistID {
			// pass the artist into the template
			err := Templates.ExecuteTemplate(w, "artist.html", artist)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}
	}
}

// ierror page handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, message string, statusCode int) {
	mu.Lock()
	defer mu.Unlock()

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
