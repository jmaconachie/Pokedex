package main

import (
	"time"

	"github.com/jmaconachie/pokedexcli/internal/pokeapi"
	"github.com/jmaconachie/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	url := "https://pokeapi.co/api/v2/location"
	cfg := &config{
		pokeapiClient:    pokeClient,
		nextLocationsURL: &url,
	}
	cache := pokecache.NewCache(5 * time.Minute)
	startRepl(cfg, cache)
}
