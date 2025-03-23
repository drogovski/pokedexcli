package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

type config struct {
	Next     string
	Previous string
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := config{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := cleanInput(scanner.Text())

		if len(userInput) == 0 {
			continue
		}
		executeCommand(userInput[0], &conf)
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
			callback:    commandMap,
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

func executeCommand(input string, conf *config) error {
	commands := getCommands()

	command, ok := commands[input]
	if ok {
		command.callback(conf)
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
