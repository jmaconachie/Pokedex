package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jmaconachie/pokedexcli/internal/pokecache"
)

func (c *Client) GetPokemon(cache *pokecache.Cache, pokemonName string) (Pokemon, error) {
	url := BaseURL + "/pokemon/" + pokemonName
	var pokemon Pokemon
	var err error

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	cache.Add(url, data)
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, err
}
