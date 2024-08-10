package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Athooh/groupie-tracker/internal/api"
)

func main() {
	err := api.InitData()
	if err != nil {
		log.Fatalf("Error initializing data: %v", err)
	}

	http.HandleFunc("/", api.IndexHandler)
	http.HandleFunc("/artists", api.ArtistsHandler)
	http.HandleFunc("/artist/", api.ArtistDetailHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server is running on port localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
