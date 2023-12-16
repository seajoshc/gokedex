package main

import "fmt"

func commandHelp(c *config) {
	fmt.Println("")
	fmt.Println("Gokedex commands are: ")
	fmt.Println("")
	commands := validCommands()
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	fmt.Println("")
}
