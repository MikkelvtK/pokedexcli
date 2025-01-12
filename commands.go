package main

import (
	"fmt"
	"os"
)

type command struct {
	name        string
	description string
	callback    func([]string, *config) error
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
		"explore": {
			name:        "explore",
			description: "Explore area for Pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon",
			callback:    commandCatch,
		},
	}
}

func commandHelp(args []string, conf *config) error {
	commands := getCommandRegistry()
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for k, c := range commands {
		fmt.Printf("%s: %s\n", k, c.description)
	}
	return nil
}

func commandExit(args []string, conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(args []string, conf *config) error {
	result, err := conf.pokeAPI.LocationAreas(conf.next)
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

func commandMapb(args []string, conf *config) error {
	if len(conf.previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}

	result, err := conf.pokeAPI.LocationAreas(conf.previous)
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

func commandExplore(args []string, conf *config) error {
	if len(args) < 2 {
		fmt.Println("no area name provided")
		return nil
	}

	fmt.Printf("Exploring %s...\n", args[1])

	result, err := conf.pokeAPI.LocationAreaEncounters(args[1])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, pokemon := range result.PokemonEncounters {
		fmt.Printf(("- %s\n"), pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(args []string, conf *config) error {
	if len(args) < 2 {
		fmt.Println("no pokemon name provided")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[1])
	return nil
}
