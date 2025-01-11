package main

import (
	"bufio"
	"fmt"
	"strings"
)

type config struct {
	next     string
	previous string
}

func run(s *bufio.Scanner) error {
	if s == nil {
		return fmt.Errorf("no scanner was provided")
	}

	commands := getCommandRegistry()
	conf := config{}

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

		if err := c.callback(&conf); err != nil {
			return err
		}
	}
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	return strings.Fields(loweredText)
}
