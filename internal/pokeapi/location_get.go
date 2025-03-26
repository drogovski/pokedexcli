package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseUrl + fmt.Sprintf("/location-area/%s", locationName)

	entry, exists := c.cache.Get(url)

	if exists {
		return createPokemonResponse(entry)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)

	return createPokemonResponse(data)
}

func createPokemonResponse(data []byte) (Location, error) {
	locationResp := Location{}
	err := json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, err
	}
	return locationResp, nil
}
