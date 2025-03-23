package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/drogovski/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		userInput := cleanInput(reader.Text())

		if len(userInput) == 0 {
			continue
		}
		err := executeCommand(userInput[0], cfg)

		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Returns next 20 locations from pokedex API",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Returns previous 20 locations from pokedex API",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func executeCommand(input string, cfg *config) error {
	commands := getCommands()

	command, exists := commands[input]
	if exists {
		err := command.callback(cfg)
		if err != nil {
			return err
		}
		return nil
	}
	fmt.Print("Unknown command")
	return nil
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}
