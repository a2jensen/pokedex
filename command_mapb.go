package main

import (
	"fmt"
	"net/url"
	"strconv"
)

func commandMapb(c *config) error {
	listResp, err := c.API.ListLocationAreas(*c.Previous)
	if err != nil {
		return fmt.Errorf("error fetching : %d", err)
	}

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