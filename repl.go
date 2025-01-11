package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type command struct {
	name        string
	description string
	callback    func() error
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
	}
}

func run(s *bufio.Scanner) error {
	if s == nil {
		return fmt.Errorf("no scanner was provided")
	}

	commands := getCommandRegistry()

	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		input := cleanInput(s.Text())

		if len(input) == 0 {
			continue
		}

		c, ok := commands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := c.callback(); err != nil {
			return err
		}
	}
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	return strings.Fields(loweredText)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	commands := getCommandRegistry()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for k, c := range commands {
		fmt.Printf("%s: %s\n", k, c.description)
	}
	return nil
}
