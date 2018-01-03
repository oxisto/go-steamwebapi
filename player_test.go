package steamwebapi

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestGetPlayerSummaries(t *testing.T) {
	apiKey := os.Getenv("STEAM_API_KEY")

	assert.NotEmpty(t, apiKey)

	response, err := GetPlayerSummaries(apiKey, []string{"76561197962272442"})

	assert.Empty(t, err)

	assert.NotEmpty(t, response)

	assert.Equal(t, 1, len(response.Response.Players))

	assert.Equal(t, "oxisto", response.Response.Players[0].PersonaName)
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
}

func TestGetPlayerAchievements(t *testing.T) {
	steamID := "76561197962272442"
	gameID := "218620"

	apiKey := os.Getenv("STEAM_API_KEY")

	assert.NotEmpty(t, apiKey)

	response, err := GetPlayerAchievements(apiKey, gameID, steamID)

	assert.Empty(t, err)

	assert.NotEmpty(t, response)

	assert.Equal(t, steamID, response.PlayerStats.SteamID)
	assert.Equal(t, "PAYDAY 2", response.PlayerStats.GameName)
	assert.NotEmpty(t, response.PlayerStats.Achievements)

	// pick the last one
	unlockedAchievement := response.GetLatestUnlockedAchievement()

	assert.NotEmpty(t, unlockedAchievement.UnlockTime)

	// see, if we can find the achievement info
	response2, err := GetSchemaForGame(apiKey, gameID)

	assert.Empty(t, err)

	assert.NotEmpty(t, response2)

	achievement := response2.Game.FindAchievement(unlockedAchievement.ApiName)

	assert.NotEmpty(t, achievement)
}
