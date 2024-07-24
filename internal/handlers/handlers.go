package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Athooh/groupie-tracker/internal/api"
)

// Template caching
var templates = template.Must(template.ParseGlob(filepath.Join("web", "templates", "*.html")))

func init() {
	if len(templates.Templates()) == 0 {
		log.Fatal("No templates found. Check your template directory and path.")
	}
}

// HomeHandler serves the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// LocationsHandler serves the locations page
func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	locations, err := api.GetLocations()
	if err != nil {
		log.Printf("Error fetching locations: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Locations []api.Location
	}{
		Locations: locations,
	}

	err = templates.ExecuteTemplate(w, "locations.html", data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// DatesHandler serves the dates page
func DatesHandler(w http.ResponseWriter, r *http.Request) {
	dates, err := api.GetDates()
	if err != nil {
		log.Printf("Error fetching dates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Dates []api.Date
	}{
		Dates: dates,
	}

	err = templates.ExecuteTemplate(w, "dates.html", data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
