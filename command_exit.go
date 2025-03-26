package main

import (
	"fmt"
	"os"
)

func commandExit(conf *config, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
