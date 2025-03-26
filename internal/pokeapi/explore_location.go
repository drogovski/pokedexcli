package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(locationName string) (RespPokemonEncounters, error) {
	url := baseUrl + fmt.Sprintf("/location-area/%s", locationName)

	entry, exists := c.cache.Get(url)

	if exists {
		return createPokemonResponse(entry)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonEncounters{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonEncounters{}, err
	}

	return createPokemonResponse(data)
}

func createPokemonResponse(data []byte) (RespPokemonEncounters, error) {
	locationResp := RespPokemonEncounters{}
	err := json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespPokemonEncounters{}, err
	}
	return locationResp, nil
}
