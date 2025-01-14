package pokeapi

import "fmt"

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

func (p *PokeAPI) LocationAreaEncounters(name string) (AreaEncountersApi, error) {
	if len(name) == 0 {
		return AreaEncountersApi{}, fmt.Errorf("no area name was provided")
	}

	url := fmt.Sprintf("%s%s/%s", baseUrl, locationAreaEndpoint, name)

	result, err := getParsedResponse[AreaEncountersApi](url, p.cache)
	if err != nil {
		return AreaEncountersApi{}, err
	}

	return result, nil
}

func (p *PokeAPI) LocationAreas(url string) (LocationApi, error) {
	if len(url) == 0 {
		url = fmt.Sprintf("%s%s", baseUrl, locationAreaEndpoint)
	}

	result, err := getParsedResponse[LocationApi](url, p.cache)
	if err != nil {
		return LocationApi{}, err
	}
	return result, nil
}
