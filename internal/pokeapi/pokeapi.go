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

func NewPokeAPI(interval time.Duration) *PokeAPI {
	return &PokeAPI{
		cache: pokecache.NewCache(interval),
	}
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
