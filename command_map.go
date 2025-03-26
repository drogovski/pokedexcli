package main

import (
	"errors"
	"fmt"

	"github.com/drogovski/pokedexcli/internal/pokeapi"
)

func commandMapf(cfg *config, args ...string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.previousLocationsURL = locationResp.Previous

	showLocations(locationResp.Results)
	return nil
}

func commandMapb(cfg *config, args ...string) error {
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

func showLocations(locations []pokeapi.ShallowLocation) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}
