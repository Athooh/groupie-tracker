package handlers

import (
	"log"
	"net/http"

	"github.com/Athooh/groupie-tracker/internal/api"
)

// ArtistsHandler serves the artists page
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := api.GetArtists()
	if err != nil {
		log.Printf("Error fetching artists: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Artists []api.Artist
	}{
		Artists: artists,
	}

	err = templates.ExecuteTemplate(w, "artists.html", data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
