package api

import (
	"fmt"
	"github.com/oxisto/go-steamwebapi/structs"
)

func GetSchemaForGame(apiKey string, gameID string) (response structs.GetSchemaForGameResponse, err error) {
	u := fmt.Sprintf("http://api.steampowered.com/ISteamUserStats/GetSchemaForGame/v2/?key=%s&appid=%s", apiKey, gameID)

	err = Request(u, &response)

	return
}

func GetUserStatsForGame(apiKey string, gameID string, steamID string) (response structs.GetUserStatsForGameResponse, err error) {
	u := fmt.Sprintf("http://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v0002/?appid=%s&key=%s&steamid=%s", gameID, apiKey, steamID)

	err = Request(u, &response)

	return
}
