package main

import "fmt"


func commandHelp() {
	fmt.Printf(
			"Welcome to %v! These are the available commands: \n",
			cliName,
	)
	fmt.Println("help    - Show available commands")
	fmt.Println("map   - Show the names of 20 location areas")
	fmt.Println("mapb   - Show previous names of 20 location areas")
	fmt.Println("clear   - Clear the terminal screen")
	fmt.Println("exit    - Closes your connection")
}