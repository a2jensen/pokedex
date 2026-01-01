package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(c *config, params ...string) error {
	if len(params) == 0 {
		fmt.Println("Pokemon name not specified")
		return nil
	}
	pokemonName := params[0]

	if _ , exists := c.Inventory[pokemonName]; exists {
		fmt.Println("Pokemon is already caught!")
		return nil
	}

	message := fmt.Sprintf("Throwing a Pokeball at %s...", pokemonName)
	fmt.Println(message)

	pokemon , err := c.API.ListPokemon(pokemonName)
	if err != nil {
		fmt.Errorf("error : %w", err)
		return nil
	}

	var catchChance int
	switch {
	case pokemon.BaseExperience < 100:
		catchChance = 80  // easy (Caterpie, Pidgey)
	case pokemon.BaseExperience < 200:
		catchChance = 50  // medium (Pikachu, starters)
	case pokemon.BaseExperience < 300:
		catchChance = 25  // hard (evolved Pokemon)
	default:
		catchChance = 10  // legendary tier
	}

	for range 3 {
		time.Sleep(1 * time.Second)
		fmt.Println("Attempting to catch.... **shake**")
		if rand.Intn(100) < catchChance {
			caughtMsg := fmt.Sprintf("%s was caught!", pokemonName)
			fmt.Println(caughtMsg)
			c.Inventory[pokemonName] = pokemon
			break
		} else {
			fmt.Println("Could not catch!")
		}
	}


	return nil
}