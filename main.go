package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	limit := 0
  offset := 0
	cfg := config{
		limit: limit,
		offset: offset,
		next:     fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=%v&offset=%v", limit, offset),
		previous: fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=%v&offset=%v", limit, offset),
	}
	commands := returnCommands(commandHelp, commandClear, func(){commandMap(&cfg)}, func(){commandMapb(&cfg)})
	// Begin the repl loop
	reader := bufio.NewScanner(os.Stdin)
	// commandHelp()
	printPrompt()
	for reader.Scan() {
			// sanitize the input
			text := cleanInput(reader.Text())
			if command, exists := commands[text]; exists {
					// Call commands callback function
					command.callback()
			} else if strings.EqualFold("exit", text) {
					// Close the program on the exit command
					return
			} else {
					// Invalid command, pass the command to the parser
					handleCmd(text)
			}
			printPrompt()
	}
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}