package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/MikkelvtK/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeAPI  *pokeapi.PokeAPI
	commands map[string]command
	scanner  *bufio.Scanner
	next     string
	previous string
}

func run(conf *config) error {
	for {
		fmt.Print("Pokedex > ")
		conf.scanner.Scan()
		input := cleanInput(conf.scanner.Text())

		if len(input) == 0 {
			continue
		}

		c, ok := conf.commands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := c.callback(conf); err != nil {
			return err
		}
	}
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	return strings.Fields(loweredText)
}
