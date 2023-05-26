package main

import (
	"os"
	"os/exec"
)

// clearScreen clears the terminal screen
func commandClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}