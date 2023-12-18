package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/seajoshc/gokedex/internal/pokeapi"
)

func main() {
	fmt.Println("📺 Booting up the Gokedex")
	scanner := bufio.NewScanner(os.Stdin)
	commands := validCommands() //commands.go

	// new pokeapi.co client for the session
	c := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	// The REPL
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

		err := command.callback(&c)

		if err != nil {
			fmt.Printf("🤯 %v", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// wraps the pokeapi client in config & state
type config struct {
	pokeapiClient             pokeapi.Client
	nextLocationAreasPage     *string
	previousLocationAreasPage *string
}

func cleanInput(input string) []string {
	lowercase := strings.TrimSpace(input)
	words := strings.Fields(lowercase)
	return words
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
			description: "Gets the next page of location areas: sections of areas, such as floors in a building or cave. Each area has its own set of possible Pokémon encounters.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Paginates back through location areas: sections of areas, such as floors in a building or cave. Each area has its own set of possible Pokémon encounters.",
			callback:    commandMapb,
		},
	}
}
