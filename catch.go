package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
)

type Pokemon struct {
	Name   string            `json:"name"`
	Height int               `json:"height"`
	Weight int               `json:"weight"`
	Stats  []StatInformation `json:"stats"`
	Types []TypeInformation  `json:"types"`
	BaseExperience int       `json:"base_experience"`
}

type TypeInformation struct {
	Type StatData   `json:"type"`
}
 

type StatInformation struct {
	Stat     StatData `json:"stat"`
	BaseStat int      `json:"base_stat"`
}

type StatData struct {
	Name string `json:"name"`
}

var caughtPokemons = make(map[string]Pokemon)

func catchCommand(pokemon string) {
	_, ok := caughtPokemons[pokemon] 
	// Check if pokemon is already catched
	if ok {
		fmt.Println("You already catched this pokemon !")
		return
	}
	
	exp := fmt.Sprintf("Throwing a Pokeball at %s....", pokemon)
	fmt.Println(exp)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemon)
	body := fetchData(url)

	var result Pokemon
	err := json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	baseExp := result.BaseExperience

	difficulty := int(baseExp) // Convert to int

	// Generate a random number between 1 and difficulty
	randomNumber := rand.Intn(difficulty) + 1
	fmt.Println(randomNumber)

	// Set the variable to true randomly with decreasing probability
	isCaught := randomNumber < 35
	if isCaught {
		caughtPokemons[pokemon] = result
		fmt.Println(caughtPokemons)
		fmt.Println(pokemon + " was caught!")
	} else {
		fmt.Println(pokemon + " escaped!")
	}
}