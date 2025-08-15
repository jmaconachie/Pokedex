package main

import (
	"errors"
	"fmt"

	"github.com/jmaconachie/pokedexcli/internal/pokeapi"
	"github.com/jmaconachie/pokedexcli/internal/pokecache"
)

func commandMapf(cfg *config, cache *pokecache.Cache) error {
	locations := pokeapi.RespShallowLocations{}
	var err error
	if cfg.nextLocationsURL == nil {
		return errors.New("no next page URL available")
	}
	if data, exists := cache.Get(*cfg.nextLocationsURL); exists {
		locations, err = pokeapi.ParseLocations(data)
		if err != nil {
			return err
		}
	} else {
		locations, err = cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL, cache)
		if err != nil {
			return err
		}
	}
	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)

	}
	return nil
}

func commandMapb(cfg *config, cache *pokecache.Cache) error {
	locations := pokeapi.RespShallowLocations{}
	var err error
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	if data, exists := cache.Get(*cfg.prevLocationsURL); exists {
		locations, err = pokeapi.ParseLocations(data)
		if err != nil {
			return err
		}
	} else {
		locations, err = cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL, cache)
		if err != nil {
			return err
		}
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
