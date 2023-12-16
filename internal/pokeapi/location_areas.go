package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// method for pokeapi.NewClient()
func (c *Client) GetLocationAreas() (LocationAreasResp, error) {
	endpoint := "/location-area"
	url := baseURL + endpoint

	res, err := http.Get(url)
	if err != nil {
		return LocationAreasResp{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
	}

	locationAreaResp := LocationAreasResp{}
	err = json.Unmarshal(body, &locationAreaResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreaResp, nil
}
