package api

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestGetSchemaForGame(t *testing.T) {
	apiKey := os.Getenv("STEAM_API_KEY")

	assert.NotEmpty(t, apiKey)

	response, err := GetSchemaForGame(apiKey, "218620")

	assert.Empty(t, err)

	assert.NotEmpty(t, response)

	assert.Equal(t, "PAYDAY 2", response.Game.GameName)

	fmt.Print(response)
}


func TestGetUserStatsForGame(t *testing.T) {
	steamID := "76561197962272442"

	apiKey := os.Getenv("STEAM_API_KEY")

	assert.NotEmpty(t, apiKey)

	response, err := GetUserStatsForGame(apiKey, "218620", steamID)

	assert.Empty(t, err)

	assert.NotEmpty(t, response)

	assert.Equal(t, steamID, response.PlayerStats.SteamID)
	assert.Equal(t, "PAYDAY 2", response.PlayerStats.GameName)
	assert.NotEmpty(t, response.PlayerStats.Achievements)
	assert.NotEmpty(t, response.PlayerStats.Stats)

	fmt.Print(response)
}
