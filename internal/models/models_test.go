package models

import (
	"testing"
)

// Test for the Artist model
func TestArtist(t *testing.T) {
	artist := Artist{
		ID:           1,
		Image:        "image_url",
		Name:         "Artist 1",
		Members:      []string{"Member 1", "Member 2"},
		CreationDate: 1990,
		FirstAlbum:   "First Album",
	}

	if artist.ID != 1 {
		t.Errorf("expected ID to be 1, got %v", artist.ID)
	}

	if artist.Image != "image_url" {
		t.Errorf("expected Image to be 'image_url', got %v", artist.Image)
	}

	if artist.Name != "Artist 1" {
		t.Errorf("expected Name to be 'Artist 1', got %v", artist.Name)
	}

	if len(artist.Members) != 2 {
		t.Errorf("expected 2 members, got %v", len(artist.Members))
	}

	if artist.CreationDate != 1990 {
		t.Errorf("expected CreationDate to be 1990, got %v", artist.CreationDate)
	}

	if artist.FirstAlbum != "First Album" {
		t.Errorf("expected FirstAlbum to be 'First Album', got %v", artist.FirstAlbum)
	}
}

// Test for the Location model
func TestLocation(t *testing.T) {
	location := Location{
		ID:        1,
		Locations: []string{"New York", "Los Angeles"},
		DatesURL:  "dates_url",
	}

	if location.ID != 1 {
		t.Errorf("expected ID to be 1, got %v", location.ID)
	}

	if len(location.Locations) != 2 {
		t.Errorf("expected 2 locations, got %v", len(location.Locations))
	}

	if location.DatesURL != "dates_url" {
		t.Errorf("expected DatesURL to be 'dates_url', got %v", location.DatesURL)
	}
}

// Test for the LocationsData model
func TestLocationsData(t *testing.T) {
	locationsData := LocationsData{
		Index: []Location{
			{ID: 1, Locations: []string{"New York"}},
			{ID: 2, Locations: []string{"Los Angeles"}},
		},
	}

	if len(locationsData.Index) != 2 {
		t.Errorf("expected 2 locations in Index, got %v", len(locationsData.Index))
	}

	if locationsData.Index[0].ID != 1 {
		t.Errorf("expected first location ID to be 1, got %v", locationsData.Index[0].ID)
	}
}

// Test for the Date model
func TestDate(t *testing.T) {
	date := Date{
		ID:    1,
		Dates: []string{"2024-01-01", "2024-01-02"},
	}

	if date.ID != 1 {
		t.Errorf("expected ID to be 1, got %v", date.ID)
	}

	if len(date.Dates) != 2 {
		t.Errorf("expected 2 dates, got %v", len(date.Dates))
	}

	if date.Dates[0] != "2024-01-01" {
		t.Errorf("expected first date to be '2024-01-01', got %v", date.Dates[0])
	}
}

// Test for the DatesData model
func TestDatesData(t *testing.T) {
	datesData := DatesData{
		Index: []Date{
			{ID: 1, Dates: []string{"2024-01-01"}},
			{ID: 2, Dates: []string{"2024-01-02"}},
		},
	}

	if len(datesData.Index) != 2 {
		t.Errorf("expected 2 dates in Index, got %v", len(datesData.Index))
	}

	if datesData.Index[0].ID != 1 {
		t.Errorf("expected first date ID to be 1, got %v", datesData.Index[0].ID)
	}
}

// Test for the Relation model
func TestRelation(t *testing.T) {
	relation := Relation{
		ID: 1,
		DatesLocations: map[string][]string{
			"2024-01-01": {"New York", "Los Angeles"},
		},
	}

	if relation.ID != 1 {
		t.Errorf("expected ID to be 1, got %v", relation.ID)
	}

	if len(relation.DatesLocations) != 1 {
		t.Errorf("expected 1 date in DatesLocations, got %v", len(relation.DatesLocations))
	}

	if locations, ok := relation.DatesLocations["2024-01-01"]; !ok {
		t.Errorf("expected date '2024-01-01' to be present")
	} else if len(locations) != 2 {
		t.Errorf("expected 2 locations for date '2024-01-01', got %v", len(locations))
	}
}

// Test for the RelationsData model
func TestRelationsData(t *testing.T) {
	relationsData := RelationsData{
		Index: []Relation{
			{
				ID: 1,
				DatesLocations: map[string][]string{
					"2024-01-01": {"New York", "Los Angeles"},
				},
			},
			{
				ID: 2,
				DatesLocations: map[string][]string{
					"2024-01-02": {"Chicago", "San Francisco"},
				},
			},
		},
	}

	if len(relationsData.Index) != 2 {
		t.Errorf("expected 2 relations in Index, got %v", len(relationsData.Index))
	}

	if relationsData.Index[0].ID != 1 {
		t.Errorf("expected first relation ID to be 1, got %v", relationsData.Index[0].ID)
	}
}
