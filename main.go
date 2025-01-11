package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/MikkelvtK/pokedexcli/internal/pokecache"
)

func main() {
	c := &config{
		commands: getCommandRegistry(),
		scanner:  bufio.NewScanner(os.Stdin),
		cache:    pokecache.NewCache(5 * time.Minute),
	}

	if err := run(c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
