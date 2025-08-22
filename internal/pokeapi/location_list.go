package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jmaconachie/pokedexcli/internal/pokecache"
)

// ListLocations -
func (c *Client) GetLocations(pageURL *string, cache *pokecache.Cache) (RespShallowLocations, error) {
	url := BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	cache.Add(url, data)
	locations, err := ParseLocations(data)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locations, nil
}

// parse Locations
func ParseLocations(data []byte) (RespShallowLocations, error) {
	locations := RespShallowLocations{}
	err := json.Unmarshal(data, &locations)
	if err != nil {
		return RespShallowLocations{}, err
	}
	return locations, nil
}
