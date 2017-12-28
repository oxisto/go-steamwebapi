package steamwebapi

import (
	"encoding/json"
	"net/http"
)

func Request(url string, response interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}

	return nil
}
