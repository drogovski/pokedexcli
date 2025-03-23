package main

import (
	"fmt"
	"os"
)

func commandExit(conf *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
