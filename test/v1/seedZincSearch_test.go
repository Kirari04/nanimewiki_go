package controllers__test

import (
	"ch/kirari/animeApi/models"
	"ch/kirari/animeApi/setups"
	"testing"
)

// TESTS

func TestSeedZincSearch_AddEntrys(t *testing.T) {
	AssertEqualBool(t, true, setups.ZincSearch_AddEntrys([]models.Anime{
		{
			Title:  "Test Entry",
			Type:   "TV",
			Status: "FINISHED",
		},
	}))
}
