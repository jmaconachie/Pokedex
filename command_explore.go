package main

import (
	"fmt"

	"github.com/jmaconachie/pokedexcli/internal/pokeapi"
	"github.com/jmaconachie/pokedexcli/internal/pokecache"
)

func commandExplore(cfg *config, cache *pokecache.Cache, args []string) error {
	var pokemon []string
	var err error
	if len(args) != 1 {
		return fmt.Errorf("incorrect amount of args")
	}
	location := args[0]
	if data, exists := cache.Get(pokeapi.BaseURL + "location-area/" + location); exists {
		pokemon, err = pokeapi.ParsePokemonList(data)
		if err != nil {
			return err
		}
	} else {
		pokemon, err = cfg.pokeapiClient.GetLocationPokemon(cache, location)
		if err != nil {
			return err
		}
	}
	fmt.Println("Exploring " + location + "\nFound Pokemon: ")
	for _, pokemonName := range pokemon {
		fmt.Println("- " + pokemonName)
	}
	return nil
}
