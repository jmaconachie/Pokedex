package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"

	"github.com/jmaconachie/pokedexcli/internal/pokeapi"
	"github.com/jmaconachie/pokedexcli/internal/pokecache"
)

func commandCatch(cfg *config, cache *pokecache.Cache, args []string) error {
	var pokemon pokeapi.Pokemon
	var err error
	if args == nil {
		fmt.Println("Please enter a pokemon to try and catch")
		return nil
	}
	pokemonName := args[0]
	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")
	if data, exists := cache.Get(pokeapi.BaseURL + "pokemon/" + pokemonName); exists {
		err = json.Unmarshal(data, &pokemon)
		if err != nil {
			return err
		}
	} else {
		pokemon, err = cfg.pokeapiClient.GetPokemon(cache, pokemonName)
		if err != nil {
			return err
		}
	}
	if pokemon.BaseExperience <= rand.IntN(255) {
		cfg.pokedex[pokemonName] = pokemon
		fmt.Println(pokemonName + " was caught!")
	} else {
		fmt.Println(pokemonName + " escaped!")
	}
	return nil
}
