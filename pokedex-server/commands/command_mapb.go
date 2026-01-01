package main

import (
	"fmt"
	"net/url"
	"strconv"
	"github.com/a2jensen/pokedexcli/internal/pokeapi"
	"encoding/json"
)

func commandMapb(c *config, param ...string) error {
	var data pokeapi.ListResp
	cacheResp, found:= c.Cache.Get(*c.Previous)
	if found {
		fmt.Println("These locations were cached!")
		parsed, _ := url.Parse(*c.Next)
		q := parsed.Query()
		offsetStr := q.Get("offset")
		offset, _ := strconv.Atoi(offsetStr)
		offset -= 20
		next := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", max(offset, 0))
		c.Next = &next

		parsed, _ = url.Parse(*c.Next)
		q = parsed.Query()
		offsetStr = q.Get("offset")
		offset, _ = strconv.Atoi(offsetStr)
		offset -= 40
		prev := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", max(offset, 0))
		c.Previous = &prev
		
		json.Unmarshal(cacheResp, &data)

		for _ , location := range data.Results {
			fmt.Println(location.Name)
		}
		return nil
	} else {
		fmt.Println("Locations not found in cache...")
	}

	listResp, err := c.API.ListLocationAreas(*c.Previous)
	if err != nil {
		return fmt.Errorf("error fetching : %d", err)
	}

	encodedCache, err := json.Marshal(listResp)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	c.Cache.Add(*c.Previous, encodedCache)


	parsed, _ := url.Parse(*c.Next)
	q := parsed.Query()
	offsetStr := q.Get("offset")
	offset, _ := strconv.Atoi(offsetStr)
	offset -= 20
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