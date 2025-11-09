package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"github.com/a2jensen/pokedexcli/internal/pokeapi"
)

func commandMap(c *config) error {
	var data pokeapi.ListResp
	cacheResp, found := c.Cache.Get(*c.Next)
	if found {
		fmt.Println("These locations were cached!!!")
		json.Unmarshal(cacheResp, &data)

		parsed, _ := url.Parse(*c.Next)
		q := parsed.Query()
		offsetStr := q.Get("offset")
		offset, _ := strconv.Atoi(offsetStr)
		offset += 20
		next := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", max(offset, 0))
		c.Next = &next

		parsed, _ = url.Parse(*c.Next)
		q = parsed.Query()
		offsetStr = q.Get("offset")
		offset, _ = strconv.Atoi(offsetStr)
		offset -= 40
		prev := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", max(offset, 0))
		c.Previous = &prev

		for _ , location := range data.Results {
			fmt.Println(location.Name)
		}
		return nil
	} else {
		fmt.Println("Locations not found in cache...")
	}

	listResp, err := c.API.ListLocationAreas(*c.Next)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	encodedCache, err := json.Marshal(listResp)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	c.Cache.Add(*c.Next, encodedCache)


	parsed, _ := url.Parse(*c.Next)
	q := parsed.Query()
	offsetStr := q.Get("offset")
	offset, _ := strconv.Atoi(offsetStr)
	offset += 20
	next := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", max(offset, 0))
	c.Next = &next

	parsed, _ = url.Parse(*c.Next)
	q = parsed.Query()
	offsetStr = q.Get("offset")
	offset, _ = strconv.Atoi(offsetStr)
	offset -= 40
	prev := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", max(offset, 0))
	c.Previous = &prev

	for _ , location := range listResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func formatURL(link string, setOffset int) (string) {
	parsed, err := url.Parse(link)
	if err != nil {
		return ""
	}
	q := parsed.Query()
	offsetStr := q.Get("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return ""
	}
	offset += setOffset
	newURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", max(offset, 0))
	return newURL
}

type ListResp struct {
    Count    int                `json:"count"`
    Next     *string            `json:"next"`     // can be null
    Previous *string            `json:"previous"` // can be null
    Results  []NamedAPIResource `json:"results"`
}

type NamedAPIResource struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}