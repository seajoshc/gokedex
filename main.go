package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Start the REPL
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := validCommands() //commands.go

	for {
		fmt.Print("gokedex 🎮 ")

		scanner.Scan() // get next input
		input := scanner.Text()
		cleanedInput := cleanInput(input)

		if len(cleanedInput) == 0 {
			continue
		}

		// command name is always first
		command, ok := commands[cleanedInput[0]]

		if !ok {
			fmt.Println("🤖 invalid command, try help")
			continue
		}

		if command.name == "exit" {
			command.callback()
			break
		}

		command.callback()
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func validCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits Gokedex.",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Prints this help message.",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Location areas are sections of areas, such as floors in a building or cave. Each area has its own set of possible Pokémon encounters.",
			callback:    commandMap,
		},
	}
}

func cleanInput(input string) []string {
	lowercase := strings.TrimSpace(input)
	words := strings.Fields(lowercase)
	return words
}
