package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	pokedex := cfg.pokedex
	if len(pokedex) == 0 {
		return errors.New("your pokedex is empty")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
