package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jmaconachie/pokedexcli/internal/pokecache"
)

func (c *Client) GetLocationPokemon(cache *pokecache.Cache, location string) ([]string, error) {
	url := BaseURL + "/location-area/" + location
	var pokemonList []string
	var err error
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemonList, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokemonList, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemonList, err
	}
	cache.Add(url, data)
	pokemonList, err = ParsePokemonList(data)
	if err != nil {
		return pokemonList, err
	}

	return pokemonList, nil
}

// parse Locations
func ParsePokemonList(data []byte) ([]string, error) {
	var LocationResponse Location
	var pokemonList []string

	err := json.Unmarshal(data, &LocationResponse)
	if err != nil {
		return pokemonList, err
	}
	for _, pokemon := range LocationResponse.PokemonEncounters {
		pokemonList = append(pokemonList, pokemon.Pokemon.Name)
	}
	return pokemonList, nil
}
