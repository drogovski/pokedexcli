package main

import (
	"errors"
	"fmt"

	"github.com/drogovski/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	showPokemonsInArea(location)

	return nil
}

func showPokemonsInArea(location pokeapi.Location) {
	fmt.Println("Found Pokemon:")
	for _, location := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", location.Pokemon.Name)
	}
}
