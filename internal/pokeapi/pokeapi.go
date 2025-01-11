package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

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

func Get[T any](url string) (T, error) {
	var result T

	res, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	return result, nil
}
