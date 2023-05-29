package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Pokemon struct {
	Name          string
	BaseExperience float64
}


var catchedPokemons = make(map[string]Pokemon)

func catchCommand(pokemon string) {
	_, ok := catchedPokemons[pokemon] 
	// Check if pokemon is already catched
	if ok {
		fmt.Println("You already catched this pokemon !")
		return
	}
	
	exp := fmt.Sprintf("Throwing a Pokeball at %s....", pokemon)
	fmt.Println(exp)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemon)
	body := fetchData(url)

	var result map[string]interface{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	baseExp := result["base_experience"].(float64) // Convert to float64

	difficulty := int(baseExp) // Convert to int

	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random number between 1 and difficulty
	randomNumber := randomGenerator.Intn(difficulty) + 1
	fmt.Println(randomNumber)

	// Set the variable to true randomly with decreasing probability
	isCatched := randomNumber > 30

	if isCatched {
		catchedPokemons[pokemon] = Pokemon{
			Name:          result["forms"].([]interface{})[0].(map[string]interface{})["name"].(string),
			BaseExperience: result["base_experience"].(float64),
		}
		fmt.Println(catchedPokemons)
	} else {
		fmt.Println(pokemon + "escaped!")
	}
}