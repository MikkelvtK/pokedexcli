package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/MikkelvtK/pokedexcli/internal/pokeapi"
)

func main() {
	c := &config{
		commands: getCommandRegistry(),
		scanner:  bufio.NewScanner(os.Stdin),
		pokeAPI:  pokeapi.NewPokeAPI(5 * time.Minute),
		pokemon:  map[string]pokeapi.Pokemon{},
	}

	if err := run(c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
