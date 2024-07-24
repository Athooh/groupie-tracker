package main

import (
	"log"
	"net/http"

	"github.com/Athooh/groupie-tracker/internal/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artists", handlers.ArtistsHandler)
	http.HandleFunc("/locations", handlers.LocationsHandler)
	http.HandleFunc("/dates", handlers.DatesHandler)

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
