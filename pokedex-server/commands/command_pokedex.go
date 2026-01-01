package main

import "fmt"


func commandPokedex(c *config, params ...string) error {
	fmt.Println("Your Pokedex:")
	for _ , pokemon := range c.Inventory {
		fmt.Printf("    - %s\n", pokemon.Name)
	}
	return nil
}