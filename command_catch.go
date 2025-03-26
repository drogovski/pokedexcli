package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you have to give pokemon name with the command")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	catchScore := rand.IntN(300)

	if pokemon.BaseExperience > catchScore {
		return errors.New("the pokemon escaped")
	}

	fmt.Printf("You have successfully caught the %s!!!\n", pokemon.Name)
	cfg.pokedex[pokemon.Name] = pokemon
	return nil
}
