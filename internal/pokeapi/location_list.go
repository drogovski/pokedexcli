package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	entry, exists := c.cache.Get(url)

	if exists {
		fmt.Println("cache is used")
		return createResponse(entry)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationResp, err := createResponse(data)
	if err == nil {
		c.cache.Add(url, data)
	}
	return locationResp, err
}

func createResponse(data []byte) (RespShallowLocations, error) {
	locationResp := RespShallowLocations{}
	err := json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	return locationResp, nil
}
