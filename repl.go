package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/a2jensen/pokedexcli/internal/pokeapi"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	prev := ""
	next := "https://pokeapi.co/api/v2/location-area?limit=20&offset=0"
	configuration := config{
		Previous :		&prev,
		Next :			&next,
		API :			pokeapi.New("https://pokeapi.co/api/v2"),
	}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&configuration)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	Previous 	*string
	Next		*string
	API			pokeapi.Client
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:		"map",
			description: "Display the Pokedex locations",
			callback: commandMap,
		},
		"mapb": {
			name: 		"mapb",
			description: "Display the previous Pokedex locations",
			callback: commandMapb,
		},
	}
}
