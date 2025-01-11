package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if err := run(scanner); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
