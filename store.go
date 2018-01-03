package steamwebapi

import (
	"fmt"
	"strings"
)

type AppDetail struct {
	Data AppData
}

type AppData struct {
	AboutTheGame string `json:"about_the_game"`
	Background   string `json:"background"`
	Categories []struct {
		Description string
		Id          int
	} `json:"categories"`
	Developers          []string
	DLCs                []int  `json:"dlc"`
	DetailedDescription string `json:"detailed_description"`
	Genres []struct {
		Description string
		Id          string
	}
	MetaCritic struct {
		Score int
		Url   string
	} `json:"metacritic"`
	HeaderImage string `json:"header_image"`
	Name        string `json:"name"`
	WebSite     string `json:"website"`
}

func GetAppDetails(appIDs [] string) (response map[string]AppDetail, err error) {
	u := fmt.Sprintf("http://store.steampowered.com/api/appdetails/?appids=%s", strings.Join(appIDs, ","))

	err = Request(u, &response)

	return
}

func (a AppData) GetGenre() (genre string) {
	var s []string

	for _, v := range a.Genres {
		s = append(s, v.Description)
	}

	return strings.Join(s, "-")
}
