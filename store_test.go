package steamwebapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAppDetails(t *testing.T) {
	games, err := GetAppDetails([]string{"218620"})

	assert.Empty(t, err)

	assert.NotEmpty(t, games)

	game := games["218620"]

	assert.NotEmpty(t, game)

	assert.Equal(t, "PAYDAY 2", game.Data.Name)
}