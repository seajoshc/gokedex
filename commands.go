package main

import "fmt"

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
	}
}

func commandExit() error {
	fmt.Println("Quitting Gokedex 👋")
	return nil
}

func commandHelp() error {
	fmt.Println("")
	fmt.Println("Gokedex commands are: ")
	fmt.Println("")
	commands := validCommands()
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	fmt.Println("")
	return nil
}
