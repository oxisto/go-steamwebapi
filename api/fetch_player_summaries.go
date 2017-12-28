package api

import (
	"encoding/json"
	"net/http"
	"github.com/oxisto/go-steamwebapi/structs"
	"fmt"
	"strings"
)

func FetchPlayerSummaries(apiKey string, steamIDs []string) ([]structs.Player, error) {
	u := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", apiKey, strings.Join(steamIDs, ","))

	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	response := structs.PlayerSummariesResponse{}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Response.Players, nil
}
