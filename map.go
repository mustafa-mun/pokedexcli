package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type config struct {
	prevOffset int
	nextOffset int
	next string
	previous string
}

func getLocationAreas(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	// get result
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

  // get results array and print the location area
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