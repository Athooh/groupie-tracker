package api

import (
	"net/http"
	"strconv"
	"strings"
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

var templateFuncs = template.FuncMap{
	"Join": strings.Join, // Register the Join function
}

var templates = template.Must(template.New("").Funcs(templateFuncs).ParseGlob("templates/*.html"))

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

func ArtistDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Get artist ID from URL and convert to integer
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Find artist by ID
	var artistDetail models.ArtistDetail
	found := false
	for _, artist := range artists {
		if artist.ID == id {
			artistDetail = models.ArtistDetail{
				Artist:    artist,
				Locations: locationsData.Index[id],
				Dates:     datesData.Index[id],
				Relations: relationsData.Index[id],
			}
			found = true
			break
		}
	}

	if !found {
		http.NotFound(w, r)
		return
	}

	// Render the artist detail template
	err = templates.ExecuteTemplate(w, "artist_detail.html", artistDetail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
