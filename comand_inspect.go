package main

import (
	"fmt"

	"github.com/jmaconachie/pokedexcli/internal/pokecache"
)

func commandInspect(cfg *config, cache *pokecache.Cache, args []string) error {
	if args == nil {
		fmt.Println("Please enter a pokemon to inspect")
		return nil
	}
	pokemonName := args[0]

	if pokemon, exists := cfg.pokedex[pokemonName]; !exists {
		fmt.Println("you have not caught this pokemon")
	} else {
		fmt.Println("Name: " + pokemon.Name)
		fmt.Println("Height: ", pokemon.Height)
		fmt.Println("Weight: ", pokemon.Weight)
		fmt.Println("Stats: ")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeInfo := range pokemon.Types {
			fmt.Println("  -", typeInfo.Type.Name)
		}
	}
	return nil
}
