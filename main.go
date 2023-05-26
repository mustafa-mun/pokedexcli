package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	fmt.Println("map   - Show the names of 20 location areas")
	fmt.Println("mapb   - Show previous names of 20 location areas")
	fmt.Println("clear   - Clear the terminal screen")
	fmt.Println("exit    - Closes your connection")
}

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

// clearScreen clears the terminal screen
func commandClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
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

// cleanInput preprocesses input to the repl
func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

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