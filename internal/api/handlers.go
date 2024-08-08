package api

import (
	"net/http"
	"text/template"

	"github.com/Athooh/groupie-tracker/internal/fetch"
	"github.com/Athooh/groupie-tracker/internal/models"
)

var (
	artists       []models.Artist
	locationsData models.LocationsData
	datesData     models.DatesData
	relationsData models.RelationsData
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func InitData() error {
	var err error
	artists, locationsData, datesData, relationsData, err = fetch.FetchAllData()
	return err
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "artists.html", artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "locations.html", locationsData.Index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DatesHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "dates.html", datesData.Index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "relations.html", relationsData.Index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
