package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/MikkelvtK/pokedexcli/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type PokeAPI struct {
	cache *pokecache.Cache
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationApi struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

func NewPokeAPI(interval time.Duration) *PokeAPI {
	return &PokeAPI{
		cache: pokecache.NewCache(interval),
	}
}

func (p *PokeAPI) LocationAreas(url string) (LocationApi, error) {
	if len(url) == 0 {
		url = fmt.Sprintf("%s/location-area", baseUrl)
	}

	var err error
	data, ok := p.cache.Get(url)
	if !ok {
		data, err = get(url)
		if err != nil {
			return LocationApi{}, err
		}

		p.cache.Add(url, data)
	}

	result := LocationApi{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return LocationApi{}, err
	}

	return result, nil
}

func get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
