package groupie_tracker

import (
	"html/template"
	"log"
)

var Templates *template.Template

func InitializeTemplates() {
	var err error
	Templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
}
