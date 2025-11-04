package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"fmt"
	"strconv"
)

func commandMap(c *config) error {
	res, err := http.Get(c.Next)
	if err != nil {
		return fmt.Errorf("error fetching : %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 || res.StatusCode < 200 {
		logMessage := fmt.Sprintf("Unsuccessful fetch %d", res.StatusCode)
		fmt.Println(logMessage)
		return nil
	}
	
	parsed, _ := url.Parse(c.Next)
	q := parsed.Query()
	offsetStr := q.Get("offset")
	offset, _ := strconv.Atoi(offsetStr)
	offset += 20
	c.Previous = c.Next
	c.Next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", offset)

	parsed, _ = url.Parse(c.Next)
	q = parsed.Query()
	offsetStr = q.Get("offset")
	offset, _ = strconv.Atoi(offsetStr)
	offset -= 40
	c.Previous = fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", max(offset, 0))
	
	var locations ListResp
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		log.Fatal(err)
	}

	for _ , location := range locations.Results {
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