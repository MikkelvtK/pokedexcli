package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/MikkelvtK/pokedexcli/internal/pokecache"
)

const (
	baseUrl              = "https://pokeapi.co/api/v2"
	locationAreaEndpoint = "/location-area"
	pokemonEndpoint      = "/pokemon"
)

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

type AreaEncountersApi struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func NewPokeAPI(interval time.Duration) *PokeAPI {
	return &PokeAPI{
		cache: pokecache.NewCache(interval),
	}
}

func (p *PokeAPI) LocationAreaEncounters(name string) (AreaEncountersApi, error) {
	if len(name) == 0 {
		return AreaEncountersApi{}, fmt.Errorf("no area name was provided")
	}

	url := fmt.Sprintf("%s%s/%s", baseUrl, locationAreaEndpoint, name)

	var err error
	data, ok := p.cache.Get(url)
	if !ok {
		data, err = get(url)
		if err != nil {
			return AreaEncountersApi{}, fmt.Errorf("error exploring area: %v", err)
		}

		p.cache.Add(url, data)
	}

	result := AreaEncountersApi{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return AreaEncountersApi{}, err
	}

	return result, nil
}

func (p *PokeAPI) LocationAreas(url string) (LocationApi, error) {
	if len(url) == 0 {
		url = fmt.Sprintf("%s%s", baseUrl, locationAreaEndpoint)
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

func getParsedResponse[T any](url string, c *pokecache.Cache) (T, error) {
	if len(url) == 0 {
		return *new(T), fmt.Errorf("no url was provided")
	}

	var err error
	data, ok := c.Get(url)
	if !ok {
		data, err = get(url)
		if err != nil {
			return *new(T), err
		}

		c.Add(url, data)
	}

	result := *new(T)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return *new(T), err
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
