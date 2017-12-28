package steamwebapi

import (
	"fmt"
	"strings"
)

func GetPlayerSummaries(apiKey string, steamIDs []string) (response GetPlayerSummariesResponse, err error) {
	u := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", apiKey, strings.Join(steamIDs, ","))

	err = Request(u, &response)

	return
}

func GetUserStatsForGame(apiKey string, gameID string, steamID string) (response GetUserStatsForGameResponse, err error) {
	u := fmt.Sprintf("http://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v0002/?appid=%s&key=%s&steamid=%s", gameID, apiKey, steamID)

	err = Request(u, &response)

	return
}

func GetPlayerAchievements(apiKey string, gameID string, steamID string) (response GetPlayerAchievementsResponse, err error) {
	u := fmt.Sprintf("http://api.steampowered.com/ISteamUserStats/GetPlayerAchievements/v0001/?appid=%s&key=%s&steamid=%s", gameID, apiKey, steamID)

	err = Request(u, &response)

	return
}
