package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"github.com/mustafa-mun/pokedexcli/internal/pokecache"
)

var exploreCache = pokecache.NewCache(time.Minute * 5) 

func exploreCommand(area string) {
	exp := fmt.Sprintf("Exploring %s....", area)
	fmt.Println(exp)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", area)
	// Check if the data exists in the cache
	if data, ok := exploreCache.Get(url); ok {
		// Use the data from the cache
		proccessExplore(data)
		return
	}

	body := fetchData(url)
	// Cache the fetched data
	pokeCache.Add(url, body)
	// Proccess the fetced data
	proccessExplore(body)
}

func proccessExplore(data []byte) {
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal(err)
	}

	// Get results array and print the pokemons
	pokemons := result["pokemon_encounters"].([]interface{})
	for _, emr := range pokemons {
		encounterMethod := emr.(map[string]interface{})["pokemon"].(map[string]interface{})
		name := encounterMethod["name"].(string)
		fmt.Println("- "+ name)
	}
}
