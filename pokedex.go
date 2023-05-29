package main

import "fmt"

func commandPokedex() {
	if len(caughtPokemons) == 0 {
		fmt.Println("You don't have any pokemons yet!")
	}
	for _, pk := range caughtPokemons {
		fmt.Println("- " + pk.Name)
	}
}