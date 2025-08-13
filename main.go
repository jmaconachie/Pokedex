package main

import (
	"time"

	"github.com/jmaconachie/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	cache := NewCache()
	startRepl(cfg)
}
