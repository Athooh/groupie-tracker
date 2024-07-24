package api

import (
	"encoding/json"
	"net/http"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

func fetchData(endpoint string, target interface{}) error {
	resp, err := http.Get(baseURL + endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

type Artist struct {
	Name       string   `json:"name,omitempty"`
	Image      string   `json:"image,omitempty"`
	Activity   int      `json:"activity,omitempty"`
	FirstAlbum string   `json:"first_album,omitempty"`
	Members    []string `json:"members,omitempty"`
}

type Location struct {
	Location []string `json:"location,omitempty"`
}

type Date struct {
	Dates []string `json:"dates,omitempty"`
}

type Relation struct {
	Artist   string   `json:"artist,omitempty"`
	Location []string `json:"location,omitempty"`
	Dates    []string `json:"dates,omitempty"`
}
