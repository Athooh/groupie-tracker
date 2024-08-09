package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"

	"github.com/Athooh/groupie-tracker/internal/models"
)

func init() {
	// Initialize templates with functions
	templates = template.Must(template.New("").Funcs(template.FuncMap{
		"Join": strings.Join,
	}).ParseGlob("../../templates/*.html"))

	// Sample data for testing
	artists = []models.Artist{
		{ID: 1, Name: "Artist 1", Members: []string{"Member 1", "Member 2"}, CreationDate: 1990, FirstAlbum: "First Album 1"},
		{ID: 2, Name: "Artist 2", Members: []string{"Member A", "Member B"}, CreationDate: 2000, FirstAlbum: "First Album 2"},
	}
	locationsData = models.LocationsData{
		Index: []models.Location{
			{ID: 1, Locations: []string{"Location 1", "Location 2"}},
			{ID: 2, Locations: []string{"Location A", "Location B"}},
		},
	}
	datesData = models.DatesData{
		Index: []models.Date{
			{ID: 1, Dates: []string{"2023-01-01", "2023-01-02"}},
			{ID: 2, Dates: []string{"2024-01-01", "2024-01-02"}},
		},
	}
	relationsData = models.RelationsData{
		Index: []models.Relation{
			{ID: 1, DatesLocations: map[string][]string{"2023-01-01": {"Location 1", "Location 2"}}},
			{ID: 2, DatesLocations: map[string][]string{"2024-01-01": {"Location A", "Location B"}}},
		},
	}
}

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestArtistsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/artists", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ArtistsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestLocationsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/locations", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LocationsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDatesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/dates", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DatesHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestRelationsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/relations", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RelationsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
