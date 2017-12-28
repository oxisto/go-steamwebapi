package steamwebapi

import (
	"fmt"
)

func GetSchemaForGame(apiKey string, gameID string) (response GetSchemaForGameResponse, err error) {
	u := fmt.Sprintf("http://api.steampowered.com/ISteamUserStats/GetSchemaForGame/v2/?key=%s&appid=%s", apiKey, gameID)

	err = Request(u, &response)

	return
}
