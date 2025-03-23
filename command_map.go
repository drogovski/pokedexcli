package main

import (
	"errors"
	"fmt"

	"github.com/drogovski/pokedexcli/internal/pokeapi"
)

func commandMapf(cfg *config) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.previousLocationsURL = locationResp.Previous

	showLocations(locationResp.Results)
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.previousLocationsURL = locationResp.Previous

	showLocations(locationResp.Results)
	return nil
}

// type pokedexResponse struct {
// 	Next     string
// 	Previous string
// 	Results  []map[string]interface{}
// }

// func commandMap(conf *config) error {
// 	var requestURL string

// 	if len(conf.Next) != 0 {
// 		requestURL = conf.Next
// 	} else {
// 		requestURL = defaultURL
// 	}

// 	res, err := http.Get(requestURL)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if res.StatusCode > 299 {
// 		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
// 		return nil
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil
// 	}

// 	var result pokedexResponse
// 	err = json.Unmarshal(body, &result)

// 	if err != nil {
// 		log.Fatal(err)
// 		return nil
// 	}

// 	newConf := config{Next: result.Next, Previous: result.Previous}
// 	*conf = newConf
// 	showLocations(result.Results)
// 	return nil
// }

func showLocations(locations []pokeapi.Location) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}
