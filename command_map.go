package main

import (
	"net/url"
	"fmt"
	"strconv"
)

func commandMap(c *config) error {
	listResp, err := c.API.ListLocationAreas(*c.Next)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	parsed, _ := url.Parse(*c.Next)
	q := parsed.Query()
	offsetStr := q.Get("offset")
	offset, _ := strconv.Atoi(offsetStr)
	offset += 20
	c.Previous = c.Next
	next := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", offset)
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