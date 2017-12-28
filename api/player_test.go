package api

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestGetPlayerSummaries(t *testing.T) {
	apiKey := os.Getenv("STEAM_API_KEY")

	assert.NotEmpty(t, apiKey)

	response, err := GetPlayerSummaries(apiKey, []string{"76561197962272442"})

	assert.Empty(t, err)

	assert.NotEmpty(t, response)

	assert.Equal(t, 1, len(response.Response.Players))

	assert.Equal(t, "oxisto", response.Response.Players[0].PersonaName)

	fmt.Print(response)
}
