package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


type config struct {
	offset int
	next string
	previous string
}

func commandMap(cfg *config) {
	// set current page to previous
	cfg.previous = cfg.next
	// Increment offset
	cfg.offset += 20
	// set next page
	cfg.next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=20&offset=%v", cfg.offset)

	// get next page
	res, err := http.Get(cfg.next)
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


func commandMapb(cfg *config) {
	// If user is not on the first page
	if cfg.offset > 0 {
		res, err := http.Get(cfg.previous)
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

		// decrease the page num
		cfg.offset -= 20
		cfg.offset -= 20

		cfg.next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=20&offset=%v",cfg.offset)


		var result map[string]interface{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			log.Fatal(err)
		}

		results := result["results"].([]interface{})
		for _, r := range results {
			location := r.(map[string]interface{})
			fmt.Println(location["name"])
		}
	} else {
		fmt.Println("You are on the first page!")
	}
	
}