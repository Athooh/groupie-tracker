package api

import (
    "testing"
)

func TestGetArtists(t *testing.T) {
    artists, err := GetArtists()
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if len(artists) == 0 {
        t.Fatalf("Expected non-zero artists, got %d", len(artists))
    }
}
