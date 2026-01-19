package main

import (
	"fmt"
	"os"
)

func main() {
	// Entry point: orchestrates the CLI startup sequence.
	// Logic should be delegated to internal packages to maintain clean separation.
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("🍩 DONAT CLI initialized.")
	return nil
}
