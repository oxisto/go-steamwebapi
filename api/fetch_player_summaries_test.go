package api

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestFetchPlayerSummaries(t *testing.T) {
	apiKey := os.Getenv("STEAM_API_KEY")

	assert.NotEmpty(t, apiKey)

	players, err := FetchPlayerSummaries(apiKey, []string{"76561197962272442"})

	assert.Empty(t, err)

	assert.NotEmpty(t, players)
}
