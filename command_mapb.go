package main

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func commandMapb(c *config) error {
	if c.Previous == "" {
		fmt.Println("You haven't even looked at the map yet!")
		return nil
	}
	res, err := http.Get(c.Previous)
	if err != nil {
		return fmt.Errorf("error fetching : %d", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return fmt.Errorf("unsuccessful Fetch : %d", res.StatusCode)
	}
	
	parsed, _ := url.Parse(c.Next)
	q := parsed.Query()
	offsetStr := q.Get("offset")
	offset, _ := strconv.Atoi(offsetStr)
	offset -= 20
	c.Next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", max(offset, 0))

	parsed, _ = url.Parse(c.Next)
	q = parsed.Query()
	offsetStr = q.Get("offset")
	offset, _ = strconv.Atoi(offsetStr)
	offset -= 40
	c.Previous = fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", max(offset, 0))


	var locations ListResp
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading in bytes %d", err)
	}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return fmt.Errorf("error : %d", err)
	}

	for _ , location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}