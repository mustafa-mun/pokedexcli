package main


import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


type config struct {
	limit int
	offset int
	next string
	previous string
}

func commandMap(cfg *config) {
	// start from page 0 (limit=0, offset=0)
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

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	// Increment limit and offset
	if cfg.limit + 20 <= int(result["count"].(float64)) {
		cfg.limit += 20
		cfg.offset += 20
	}

	cfg.next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=%v&offset=%v", cfg.limit, cfg.offset)

	results := result["results"].([]interface{})
	for _, r := range results {
		location := r.(map[string]interface{})
		fmt.Println(location["name"])
	}
}


func commandMapb(cfg *config) {
	// If user is not on the first page
	if cfg.limit > 20 {
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

		var result map[string]interface{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			log.Fatal(err)
		}

		// Substract limit and offset
			cfg.limit -= 20
			cfg.offset -= 20

		cfg.previous = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=%v&offset=%v", cfg.limit, cfg.offset)

		results := result["results"].([]interface{})
		for _, r := range results {
			location := r.(map[string]interface{})
			fmt.Println(location["name"])
		}
	} else {
		fmt.Println("You are on the first page!")
	}
	
}