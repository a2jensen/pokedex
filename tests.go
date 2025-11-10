package main

import (
	"testing"
	"time"
	"fmt"
	"github.com/a2jensen/pokedexcli/internal/pokecache"
	"encoding/json"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	data := map[string]interface{}{
	"count": 1089,
	"next":  "https://pokeapi.co/api/v2/location-area?offset=0&limit=3",
	"previous": nil,
	"results": []map[string]string{
		{"name": "canalave-city-area", "url": "https://pokeapi.co/api/v2/location-area/1/"},
		{"name": "eterna-city-area", "url": "https://pokeapi.co/api/v2/location-area/2/"},
		{"name": "pastoria-city-area", "url": "https://pokeapi.co/api/v2/location-area/3/"},
	},
	}
	dataMarshaled, _ := json.Marshal(data)
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area?offset=0&limit=3",
			val: []byte(dataMarshaled),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	data := map[string]interface{}{
	"count": 1089,
	"next":  "https://pokeapi.co/api/v2/location-area?offset=0&limit=3",
	"previous": nil,
	"results": []map[string]string{
		{"name": "canalave-city-area", "url": "https://pokeapi.co/api/v2/location-area/1/"},
		{"name": "eterna-city-area", "url": "https://pokeapi.co/api/v2/location-area/2/"},
		{"name": "pastoria-city-area", "url": "https://pokeapi.co/api/v2/location-area/3/"},
	},
	}
	dataMarshaled, _ := json.Marshal(data)
	cache.Add("https://pokeapi.co/api/v2/location-area?offset=0&limit=3", []byte(dataMarshaled))

	_, ok := cache.Get("https://pokeapi.co/api/v2/location-area?offset=0&limit=3")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://pokeapi.co/api/v2/location-area?offset=0&limit=3")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
