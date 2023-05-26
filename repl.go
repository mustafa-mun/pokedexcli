package main

import (
	"fmt"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

// name used in the repl prompts
var cliName string = "pokedex"
 
// display the repl prompt at the start of each loop
func printPrompt() {
    fmt.Print(cliName, "> ")
}
 
//  inform the user about invalid commands
func printUnknown(text string) {
	fmt.Println(text, ": command not found")
}

// attempt to recover from a bad command
func handleInvalidCmd(text string) {
	defer printUnknown(text)
}

// parse the given commands
func handleCmd(text string) {
	handleInvalidCmd(text)
}

//  preprocesses input to the repl
func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func returnCommands(commandHelp, commandClear, commandMap, commandMapb func())(map[string]cliCommand) {
	return map[string]cliCommand{
    "help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    },
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
	},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous names of 20 location areas in the Pokemon world",
			callback:    commandMapb,
	},
		"clear" :{
				name: "clear",
				description: "Clears the terminal",
				callback: commandClear,
		},
}
}