package api

import (
	"encoding/json"
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

func RenderError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	errDetail := models.ErrorDetail{
		Title:   http.StatusText(status),
		Message: message,
	}
	err := templates.ExecuteTemplate(w, strconv.Itoa(status)+".html", errDetail)
	if err != nil {
		http.Error(w, "An error occurred - Bad request", http.StatusInternalServerError)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		RenderError(w, http.StatusInternalServerError, err.Error())
	}
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "artists.html", artists)
	if err != nil {
		RenderError(w, http.StatusInternalServerError, err.Error())
	}
}

func ArtistDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Get artist ID from URL and convert to integer
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RenderError(w, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	// Adjust for zero-based index
	index := id - 1

	// Ensure the index is within the valid range of your data
	if index < 0 || index >= len(artists) {
		RenderError(w, http.StatusNotFound, "Artist not found")
		return
	}

	// Find artist by ID
	artistDetail := models.ArtistDetail{
		Artist:    artists[index],
		Locations: locationsData.Index[index],
		Dates:     datesData.Index[index],
		Relations: relationsData.Index[index],
	}

	// Render the artist detail template
	err = templates.ExecuteTemplate(w, "artist_detail.html", artistDetail)
	if err != nil {
		RenderError(w, http.StatusInternalServerError, err.Error())
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))
	var results []map[string]string

	if query != "" {
		for _, artist := range artists {
			types := artist.SearchResultType(query)
			for _, resultType := range types {
				results = append(results, map[string]string{
					"name": resultType,
					"id":   strconv.Itoa(artist.ID),
				})
			}
		}
	}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	}

