package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/seajoshc/gokedex/internal/pokeapi"
)

// Start the REPL
func main() {
	fmt.Println("📺 Booting up the Gokedex")
	scanner := bufio.NewScanner(os.Stdin)
	commands := validCommands() //commands.go

	c := config{
		pokeapiClient: pokeapi.NewClient(),
		nextPage:      nil,
		previousPage:  nil,
	}

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
			command.callback(&c)
			break
		}

		command.callback(&c)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config)
}

type config struct {
	pokeapiClient pokeapi.Client
	nextPage      *string
	previousPage  *string
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
