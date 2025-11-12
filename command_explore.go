package main

import (
	"fmt"
	"encoding/json"
)

func commandExplore(c *config, param ...string) error {
	if len(param) == 0 {
		fmt.Println("Specify a location")
		return nil
	}

	fmt.Println("Finding the pokemon at this location...")
	url := "https://pokeapi.co/api/v2/location-area/" + param[0]

	pokemonCached, found := c.Cache.Get(url)
	if found {
		var pokemonDecoded []string
		err := json.Unmarshal(pokemonCached, &pokemonDecoded)
		if err != nil {
			return fmt.Errorf("error unmarshaling cached pokemon")
		}

		for _ , name := range pokemonDecoded {
			fmt.Println(name)
		}
		return nil
	}

	pokemon, err := c.API.ListLocationPokemon(url)
	if err != nil {
		return fmt.Errorf("error fetching pokemon: %w", err)
	}

	pokemonEncoded, err := json.Marshal(pokemon)
	if err != nil {
		return fmt.Errorf("error caching pokemon: %w", err)
	}
	c.Cache.Add(url,pokemonEncoded)

	for _ , name := range pokemon {
		fmt.Println(name)
	}

	return nil
}
