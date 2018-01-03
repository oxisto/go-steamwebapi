package steamwebapi

import "strings"

type GetSchemaForGameResponse struct {
	Game Game `json:"game"`
}

type GetUserStatsForGameResponse struct {
	PlayerStats struct {
		SteamID  string `json:"steamID"`
		GameName string `json:"gameName"`
		Achievements []struct {
			Name     string `json:"name"`
			Achieved int    `json:"achieved"`
		} `json:"achievements"`
		Stats []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"stats"`
	} `json:"playerstats"`
}

type GetPlayerAchievementsResponse struct {
	PlayerStats struct {
		SteamID      string                `json:"steamID"`
		GameName     string                `json:"gameName"`
		Achievements []UnlockedAchievement `json:"achievements"`
	}
}

type GetPlayerSummariesResponse struct {
	Response struct {
		Players []Player `json:"players"`
	} `json:"response"`
}

type Game struct {
	GameName    string `json:"gameName"`
	GameVersion string `json:"gameVersion"`

	AvailableGameStats struct {
		Achievements []Achievement `json:"achievements"`
		Stats        []Stat        `json:"stats"`
	}
}

type Achievement struct {
	Name         string `json:"name"`
	DefaultValue int    `json:"defaultvalue"`
	DisplayName  string `json:"displayName"`
	Hidden       int    `json:"hidden"`
	Description  string `json:"description"`
	Icon         string `json:"icon"`
	IconGray     string `json:"icongray"`
}

type UnlockedAchievement struct {
	ApiName    string `json:"apiname"`
	Achieved   int    `json:"achieved"`
	UnlockTime int    `json:"unlocktime"`
}

type Stat struct {
	Name         string `json:"name"`
	DefaultValue int    `json:"defaultvalue"`
	DisplayName  string `json:"displayName"`
}

type Player struct {
	SteamId                  string `json:"steamid"`
	PersonaName              string `json:"personaname"`
	ProfileUrl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatarmedium"`
	AvatarFull               string `json:"avatarfull"`
	PersonaState             int    `json:"personastate"`
	CommunityVisibilityState int    `json:"communityvisibilitystate"`
	ProfileState             int    `json:"profilestate"`
	LastLogOff               int    `json:"lastlogoff"`
	CommentPermission        bool   `json:"commentpermission"`
	RealName                 string `json:"realname"`
	PrimaryClanId            string `json:"primaryclanid"`
	TimeCreated              int    `json:"timecreated"`
	GameId                   string `json:"gameid"`
	GameServerIP             string `json:"gameserverip"`
	GameExtraInfo            string `json:"gameextrainfo"`
	CityId                   int    `json:"cityid"`
	LocationCountryCode      string `json:"loccountrycode"`
	LocationStateCode        string `json:"locstatecode"`
	LocationCityId           int    `json:"loccityid"`
}

func (g Game) FindAchievement(name string) *Achievement {
	for _, achievement := range g.AvailableGameStats.Achievements {
		if achievement.Name == name {
			return &achievement
		}
	}

	return nil
}

func (r GetUserStatsForGameResponse) GetStat(name string) *struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
} {
	for _, stat := range r.PlayerStats.Stats {
		if stat.Name == name {
			return &stat
		}
	}

	return nil
}

func (r GetUserStatsForGameResponse) GetNumStatsWithPrefix(prefix string) (i int) {
	for _, stat := range r.PlayerStats.Stats {
		if strings.HasPrefix(stat.Name, prefix) {
			i++
		}
	}

	return
}

func (r GetPlayerAchievementsResponse) GetLatestUnlockedAchievement() (lastAchievement UnlockedAchievement) {
	for _, achievement := range r.PlayerStats.Achievements {
		if achievement.UnlockTime > lastAchievement.UnlockTime {
			lastAchievement = achievement
		}
	}

	return
}
