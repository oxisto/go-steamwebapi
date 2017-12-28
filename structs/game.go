package structs

type GetSchemaForGameResponse struct {
	Game Game `json:game`
}

type GetUserStatsForGameResponse struct {
	PlayerStats struct {
		SteamID  string `json:steamID`
		GameName string `json:gameName`
		Achievements []struct {
			Name     string `json:name`
			Achieved int    `json:achieved`
		} `json:achievements`
		Stats []struct {
			Name     string `json:name`
			Value int    `json:achieved`
		} `json:stats`
	} `json:playerstats`
}

type Game struct {
	GameName    string `json:gameName`
	GameVersion string `json:gameVersion`

	/*GameAvailableStats struct {
		Achievements []Achievement `json:achievements`
		Stats        []Stat        `json:stats`
	} `json:availableGameStats`*/
}

type Achievement struct {
	Name         string `json:name`
	DefaultValue int    `json:defaultvalue`
	DisplayName  string `json:displayName`
	Hidden       int    `json:hidden`
	Description  string `json:description`
	Icon         string `json:icon`
	IconGray     string `json:icongray`
}

type Stat struct {
	Name         string `json:name`
	DefaultValue int    `json:defaultvalue`
	DisplayName  string `json:displayName`
}
