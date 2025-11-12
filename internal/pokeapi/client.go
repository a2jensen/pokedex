package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client interface {
	ListLocationAreas(url string) (ListResp,error)
	ListLocationPokemon(location string) ([]string, error)
}


type HTTPClient struct {
	base 	string
	h 		*http.Client	
}

func New(base string) *HTTPClient {
	return &HTTPClient {
		base : base,
		h: http.DefaultClient,
	}
}

func (c *HTTPClient) ListLocationAreas(url string) (ListResp, error) {
	if url == "" {
		url = c.base + "/location-area"
	}

	res, err := c.h.Get(url)
	if err != nil {
		return ListResp{}, err
	}

	defer res.Body.Close()
	var LocationAreas ListResp

	LocationAreas, err = Decode[ListResp](res)
	if err != nil {
		return ListResp{}, err
	}

	return LocationAreas, nil
}

func (c *HTTPClient) ListLocationPokemon(url string) ([]string , error) {
	fmt.Println(url)

	res, err := c.h.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var locationInfo LocationInfo
	var pokemonNames []string
	
	locationInfo, err = Decode[LocationInfo](res)
	if err != nil {
		return nil, err
	}

	for _, encounter := range locationInfo.PokemonEncounters {
		pokemonNames = append(pokemonNames, encounter.Pokemon.Name)
	}

	return pokemonNames, nil
}

// general function used to decode any type(locations, pokemon, etc.)
func Decode[T any](res *http.Response) (T, error) {
	var result T
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		var zero T
		return zero, err
	}
	return result, nil
}