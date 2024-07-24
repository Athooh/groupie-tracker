package api

import (
	"log"
)

func GetArtists() ([]Artist, error) {
	var artists []Artist
	err := fetchData("/artists", &artists)
	if err != nil {
		log.Printf("Error fetching artists: %v", err)
		return nil, err
	}
	return artists, nil
}
