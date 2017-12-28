package api

import (
	"github.com/oxisto/go-steamwebapi/structs"
	"fmt"
	"strings"
)

func GetPlayerSummaries(apiKey string, steamIDs []string) (response structs.GetPlayerSummariesResponse, err error) {
	u := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", apiKey, strings.Join(steamIDs, ","))

	err = Request(u, &response)

	return
}
