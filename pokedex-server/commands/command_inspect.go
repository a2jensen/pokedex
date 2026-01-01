package main

import "fmt"

func commandInspect(c *config, params ...string) error {
	if len(params) == 0 {
		fmt.Println("Please specify a pokemon name")
		return nil
	}
	pokemonName := params[0]
	if _ , exists := c.Inventory[pokemonName]; !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	pokemon := c.Inventory[pokemonName]
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

// how it should be formatted
/**
Name: pidgey
Height: 3
Weight: 18
Stats:
  -hp: 40
  -attack: 45
  -defense: 40
  -special-attack: 35
  -special-defense: 35
  -speed: 56
Types:
  - normal
  - flying
*/