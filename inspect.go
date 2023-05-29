package main

import "fmt"

func inspectCommand(pokemon string) {
	pk, ok := caughtPokemons[pokemon]

	if !ok{
		fmt.Println("you have not caught that pokemon")
		return
	}

	fmt.Printf("Name: %s\n", pk.Name)
	fmt.Printf("Height: %d\n", pk.Height)
	fmt.Printf("Weight: %d\n", pk.Weight)
	fmt.Println("Stats:")
	for _, stat := range pk.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pk.Types {
		fmt.Printf("  - %s\n", typ.Type.Name)
	}

}