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
	callback    func(conf *config, args ...string) error
}

type config struct {
	pokeapiClient        *pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
	pokedex              map[string]pokeapi.Pokemon
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

		commandName := userInput[0]
		args := []string{}
		if len(userInput) > 1 {
			args = userInput[1:]
		}

		err := executeCommand(commandName, cfg, args...)

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
			name:        "explore <location>",
			description: "Shows all pokemons that can be encountered in provided area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Tries to catch pokemon into pokeball and place in the pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Shows information about selected pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows all caught pokemons",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func executeCommand(commandName string, cfg *config, args ...string) error {
	commands := getCommands()

	command, exists := commands[commandName]
	if exists {
		err := command.callback(cfg, args...)
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
