package main

import (
	"fmt"

	"github.com/drogovski/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, param string) error {
	fmt.Println("Exploring pastoria-city-area...")
	pokemonResp, err := cfg.pokeapiClient.ExploreLocation(param)
	if err != nil {
		return err
	}
	showPokemonsInArea(pokemonResp.PokemonEncounters)

	return err
}

func showPokemonsInArea(encounters []pokeapi.PokemonEncounter) {
	fmt.Println("Found Pokemon:")
	for _, encounter := range encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
}
