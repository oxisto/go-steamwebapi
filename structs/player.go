package structs

// TODO: enum for persona state

type GetPlayerSummariesResponse struct {
	Response struct {
		Players []Player `json:players`
	} `json:response`
}

type Player struct {
	SteamId                  string `json:steamid`
	PersonaName              string `json:personaname`
	ProfileUrl               string `json:profileurl`
	Avatar                   string `json:avatar`
	AvaterMedium             string `json:avatarmedium`
	AvatarFull               string `json:avatarfull`
	PersonaState             int    `json:personastate`
	CommunityVisibilityState int    `json:communityvisibilitystate`
	ProfileState             int    `json:profilestate`
	LastLogOff               int    `json:lastlogoff`
	CommentPermission        bool   `json:commentpermission`
	RealName                 string `json:realname`
	PrimaryClanId            string `json:primaryclanid`
	TimeCreated              int    `json:timecreated`
	GameId                   string `json:gameid`
	GameServerIP             string `json:gameserverip`
	GameExtraInfo            string `json:gameextrainfo`
	CityId                   int    `json:cityid`
	LocationCountryCode      string `json:loccountrycode`
	LocationStateCode        string `json:locstatecode`
	LocationCityId           int    `json:loccityid`
}
