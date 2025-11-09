package main

import (
	"github.com/a2jensen/pokedexcli/internal/pokeapi"
	"github.com/a2jensen/pokedexcli/internal/pokecache"

)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	Previous 	*string
	Next		*string
	API			pokeapi.Client
	Cache		pokecache.PokeCache
}