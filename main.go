package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)
type cliCommand struct {
	name        string
	description string
	callback    func()
}

// cliName is the name used in the repl prompts
var cliName string = "pokedex"
 
// printPrompt displays the repl prompt at the start of each loop
func printPrompt() {
    fmt.Print(cliName, "> ")
}
 
func commandHelp() {
	fmt.Printf(
			"Welcome to %v! These are the available commands: \n",
			cliName,
	)
	fmt.Println("help    - Show available commands")
	fmt.Println("clear   - Clear the terminal screen")
	fmt.Println("exit    - Closes your connection")
}


// clearScreen clears the terminal screen
func commandClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func returnCommands(commandHelp, commandClear func())(map[string]cliCommand) {
	return map[string]cliCommand{
    "help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    },
		"clear" :{
				name: "clear",
				description: "Clears the terminal",
				callback: commandClear,
		},
}
}

// printUnkown informs the user about invalid commands
func printUnknown(text string) {
	fmt.Println(text, ": command not found")
}

// handleInvalidCmd attempts to recover from a bad command
func handleInvalidCmd(text string) {
	defer printUnknown(text)
}

// handleCmd parses the given commands
func handleCmd(text string) {
	handleInvalidCmd(text)
}

// cleanInput preprocesses input to the db repl
func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func main() {
	commands := returnCommands(commandHelp, commandClear)
	// Begin the repl loop
	reader := bufio.NewScanner(os.Stdin)
	commandHelp()
	printPrompt()
	for reader.Scan() {
			text := cleanInput(reader.Text())
			if command, exists := commands[text]; exists {
					// Call a hardcoded function
					command.callback()
			} else if strings.EqualFold(".exit", text) {
					// Close the program on the exit command
					return
			} else {
					// Pass the command to the parser
					handleCmd(text)
			}
			printPrompt()
	}
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}