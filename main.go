package main

import (
	"time"

	"github.com/drogovski/pokedexcli/internal/pokeapi"
	"github.com/drogovski/pokedexcli/internal/pokecache"
)

func main() {
	pokeCache := pokecache.NewCache(5 * time.Second)
	pokeClient := pokeapi.NewClient(pokeCache, 5*time.Second)
	cfg := &config{
		pokeapiClient: &pokeClient,
	}
	startRepl(cfg)
}
