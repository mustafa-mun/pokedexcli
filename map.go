package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"github.com/mustafa-mun/pokedexcli/internal/pokecache"
)

type config struct {
	prevOffset int
	nextOffset int
	next string
	previous string
}

var pokeCache = pokecache.NewCache(time.Minute * 5) 

func getLocationAreas(url string) {
	// Check if the data exists in the cache
	if data, ok := pokeCache.Get(url); ok {
		// Use the data from the cache
		proccessAreas(data)
		return
	}

	body := fetchData(url)
	// Cache the fetched data
	pokeCache.Add(url, body)
	// Proccess the fetced data
	proccessAreas(body)
}

func fetchData(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	return body
}

func proccessAreas(data []byte) {
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal(err)
	}

	// Get results array and print the location area
	results := result["results"].([]interface{})
	for _, r := range results {
		location := r.(map[string]interface{})
		fmt.Println(location["name"])
	}
}


func commandMap(cfg *config) {
	// set current page to previous
	cfg.prevOffset = cfg.nextOffset
	cfg.previous = cfg.next
	// Increment offset
	cfg.nextOffset += 20
	// set next page
	cfg.next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=20&offset=%v", cfg.nextOffset)
	// get next page
	getLocationAreas(cfg.next)
}

func commandMapb(cfg *config) {
	// If user is not on the first page
	if cfg.prevOffset >= 0 {
		getLocationAreas(cfg.previous)
		// decrease the page num
		cfg.nextOffset -= 20
		cfg.prevOffset -= 20
		// set prev and next to decreased offsets
		cfg.previous = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=20&offset=%v",cfg.prevOffset)
		cfg.next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=20&offset=%v",cfg.nextOffset)
	} else {
		fmt.Println("You are on the first page!")
	}
}