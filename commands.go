package main

import (
	"fmt"
	"os"

	"github.com/MikkelvtK/pokedexcli/internal/pokeapi"
)

const baseUrl = "https://pokeapi.co/api/v2"

type command struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommandRegistry() map[string]command {
	return map[string]command{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 locations",
			callback:    commandMapb,
		},
	}
}

func commandHelp(conf *config) error {
	commands := getCommandRegistry()
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for k, c := range commands {
		fmt.Printf("%s: %s\n", k, c.description)
	}
	return nil
}

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(conf *config) error {
	if len(conf.next) == 0 {
		conf.next = fmt.Sprintf("%s/location", baseUrl)
	}

	result, err := pokeapi.Get[pokeapi.LocationApi](conf.next)
	if err != nil {
		return err
	}

	for _, loc := range result.Results {
		fmt.Println(loc.Name)
	}

	conf.previous = result.Previous
	conf.next = result.Next
	return nil
}

func commandMapb(conf *config) error {
	if len(conf.previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}

	result, err := pokeapi.Get[pokeapi.LocationApi](conf.previous)
	if err != nil {
		return err
	}

	for _, loc := range result.Results {
		fmt.Println(loc.Name)
	}

	conf.next = result.Next
	conf.previous = result.Previous
	return nil
}
