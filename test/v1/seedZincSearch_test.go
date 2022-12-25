package controllers__test

import (
	"ch/kirari/animeApi/models"
	"ch/kirari/animeApi/setups"
	"testing"
)

// TESTS

func TestSeedZincSearch_AddEntry(t *testing.T) {
	AssertEqualBool(t, true, setups.SeedZincSearch_AddEntry(models.Anime{
		Title:  "Test Entry",
		Type:   "TV",
		Status: "FINISHED",
	}))
}
