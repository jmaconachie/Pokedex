package main

import (
	"fmt"
	"os"

	"github.com/jmaconachie/pokedexcli/internal/pokecache"
)

func commandExit(cfg *config, cache *pokecache.Cache, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
