package pokeapi

import (
	"encoding/json"
	"fmt"
)

type PokemonApi struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

func (p *PokeAPI) Pokemon(name string) (PokemonApi, error) {
	if len(name) == 0 {
		return PokemonApi{}, fmt.Errorf("no pokemon name was provided")
	}

	url := fmt.Sprintf("%s%s/%s", baseUrl, pokemonEndpoint, name)

	var err error
	data, ok := p.cache.Get(url)
	if !ok {
		data, err = get(url)
		if err != nil {
			return PokemonApi{}, fmt.Errorf("error fetching pokemon: %v", err)
		}

		p.cache.Add(url, data)
	}

	result := PokemonApi{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return PokemonApi{}, err
	}

	return result, nil
}
