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
	callback    func(conf *config, param string) error
}

type config struct {
	pokeapiClient        *pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		userInput := cleanInput(reader.Text())

		var err error
		if len(userInput) == 0 {
			continue
		} else if len(userInput) == 1 {
			err = executeCommand(userInput[0], "", cfg)
		} else {
			err = executeCommand(userInput[0], userInput[1], cfg)
		}

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
			description: "Shows next 20 locations from pokedex API",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows previous 20 locations from pokedex API",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows all pokemons that can be encountered in provided area",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func executeCommand(input, param string, cfg *config) error {
	commands := getCommands()

	command, exists := commands[input]
	if exists {
		err := command.callback(cfg, param)
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
