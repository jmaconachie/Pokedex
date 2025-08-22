package main

import (
	"fmt"

	"github.com/jmaconachie/pokedexcli/internal/pokecache"
)

func commandPokedex(cfg *config, cache *pokecache.Cache, args []string) error {
	fmt.Println("Your Pokemon:")
	for pokemon := range cfg.pokedex {
		fmt.Println(" - ", cfg.pokedex[pokemon].Name)
	}
	return nil
}
