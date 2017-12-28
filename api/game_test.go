package main

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestGetSchemaForGame(t *testing.T) {
	apiKey := os.Getenv("STEAM_API_KEY")

	assert.NotEmpty(t, apiKey)

	response, err := GetSchemaForGame(apiKey, "218620")

	assert.Empty(t, err)

	assert.NotEmpty(t, response)

	assert.Equal(t, "PAYDAY 2", response.Game.GameName)
}

