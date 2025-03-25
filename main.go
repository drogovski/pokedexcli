package main

import (
	"time"

	"github.com/drogovski/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Minute, 5*time.Second)
	cfg := &config{
		pokeapiClient: &pokeClient,
	}
	startRepl(cfg)
}
