package models

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesURL  string   `json:"dates"`
}

type LocationsData struct {
	Index []Location `json:"index"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DatesData struct {
	Index []Date `json:"index"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationsData struct {
	Index []Relation `json:"index"`
}
