package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
)

type Pokemon struct {
	name            string
	baseExperience  float64
	abilities      []string
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

	var result map[string]interface{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	baseExp := result["base_experience"].(float64) // Convert to float64

	difficulty := int(baseExp) // Convert to int

	// Generate a random number between 1 and difficulty
	randomNumber := rand.Intn(difficulty) + 1
	fmt.Println(randomNumber)

	// Set the variable to true randomly with decreasing probability
	isCaught := randomNumber < 30
	if isCaught {
		pokemonAbilities := result["abilities"].([]interface{})
		abilities := make([]string, len(pokemonAbilities))

		for i, ability := range pokemonAbilities {
			abilities[i] = ability.(map[string]interface{})["ability"].(map[string]interface{})["name"].(string)
		}

		caughtPokemons[pokemon] = Pokemon{
			name:           result["forms"].([]interface{})[0].(map[string]interface{})["name"].(string),
			baseExperience: result["base_experience"].(float64),
			abilities:      abilities,
		}

		fmt.Println(caughtPokemons)
		fmt.Println(pokemon + " was caught!")
	} else {
		fmt.Println(pokemon + " escaped!")
	}
}