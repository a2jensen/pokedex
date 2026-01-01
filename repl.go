package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/a2jensen/pokedexcli/internal/pokeapi"
	"github.com/a2jensen/pokedexcli/internal/pokecache"

)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	prev := ""
	next := "https://pokeapi.co/api/v2/location-area?limit=20&offset=0"
	inventory := map[string]pokeapi.Pokemon{}
	configuration := config{
		Previous :		&prev,
		Next :			&next,
		API :			pokeapi.New("https://pokeapi.co/api/v2"),
		Cache :			pokecache.NewCache(10 * time.Second),
		Inventory :		inventory,
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
			err := command.callback(&configuration, words[1:]...)
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
		"explore": {
			name:		"explore",
			description:	"Find pokemon at a specific location",
			callback: commandExplore,
		},
		"catch": {
			name:		"catch",
			description:	"Attempt to catch a pokemon",
			callback:	commandCatch,
		},
		"inspect": {
			name:		"inspect",
			description:	"display a caught pokemon's information",
			callback:	commandInspect,
		},
		"pokedex": {
			name:		"pokedex",
			description:	"display all pokemon captured",
			callback:	commandPokedex,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, param ...string) error
}

type config struct {
	Previous 	*string
	Next		*string
	API			pokeapi.Client
	Cache		pokecache.PokeCache
	Inventory	map[string]pokeapi.Pokemon
}
