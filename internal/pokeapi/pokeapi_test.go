package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MikkelvtK/pokedexcli/internal/pokecache"
)

func TestGet(t *testing.T) {
	ts := initTestServer("hello world", t)
	defer ts.Close()

	res, err := get(ts.URL)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	json, _ := json.Marshal("hello world")

	if string(res) != string(json) {
		t.Errorf("response did not match. expected: hello world, got: %s", string(res))
	}
}

func TestGetParsedResponseFromRequest(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "hello-world",
			expected: "hello-world",
		},
		{
			input:    "squirtle-charmander-bulbasaur",
			expected: "squirtle-charmander-bulbasaur",
		},
		{
			input:    "filled-pokedex",
			expected: "filled-pokedex",
		},
	}

	cache := pokecache.NewCache(999 * time.Minute)
	for _, c := range cases {
		ts := initTestServer(c.expected, t)
		defer ts.Close()

		res, err := getParsedResponse[string](ts.URL, cache)
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}

		if res != c.expected {
			t.Errorf("response does not match expected. expected: %s, got: %s", c.expected, res)
		}
	}
}

func TestGetParsedResponseFromCache(t *testing.T) {
	cases := []struct {
		cacheInput  string
		serverInput string
	}{
		{
			cacheInput:  "hello world from the cache",
			serverInput: "hello world from the server",
		},
		{
			cacheInput:  "squirtle charmander bulbasaur",
			serverInput: "chikorita cyndaquil totodile",
		},
		{
			cacheInput:  "filled pokedex",
			serverInput: "empty pokedex",
		},
	}

	cache := pokecache.NewCache(999 * time.Minute)
	for _, c := range cases {
		ts := initTestServer(c.serverInput, t)
		input, _ := json.Marshal(c.cacheInput)
		cache.Add(ts.URL, input)

		res, err := getParsedResponse[string](ts.URL, cache)
		if err != nil {
			t.Errorf("got error: %v", err)
		}

		if res != c.cacheInput {
			t.Errorf("expected %s from cache, got: %s", c.cacheInput, res)
		}
	}
}

func initTestServer(res string, t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json, err := json.Marshal(res)
		if err != nil {
			t.Errorf("expected no errors, got: %v", err)
		}
		fmt.Fprintf(w, "%s", json)
	}))
}
