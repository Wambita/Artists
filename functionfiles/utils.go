package groupie_tracker

import (
	"html/template"
	"log"
	"sync"
)

var (
	Templates *template.Template
	once      sync.Once //ensures initialization only happens once
)

// InitializeTemplates initializes the template variable, parsing all HTML files in the templates directory.
func InitializeTemplates() {
	once.Do(func() {
		var err error
		Templates, err = template.ParseGlob("templates/*.html")
		if err != nil {
			log.Fatalf("Error parsing templates: %v", err)
		}

	})
}
