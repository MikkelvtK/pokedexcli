package pokeapi

import (
	"encoding/json"
	"fmt"
)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Order          int    `json:"order"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	}
}

func (p *PokeAPI) Pokemon(name string) (Pokemon, error) {
	if len(name) == 0 {
		return Pokemon{}, fmt.Errorf("no pokemon name was provided")
	}

	url := fmt.Sprintf("%s%s/%s", baseUrl, pokemonEndpoint, name)

	var err error
	data, ok := p.cache.Get(url)
	if !ok {
		data, err = get(url)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error fetching pokemon: %v", err)
		}

		p.cache.Add(url, data)
	}

	result := Pokemon{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return Pokemon{}, err
	}

	return result, nil
}
