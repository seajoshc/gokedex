package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := validCommands()

	for {
		fmt.Print("gokedex 🎮 ")

		scanner.Scan() // get next input
		input := scanner.Text()
		input = strings.TrimSpace(input)

		command, ok := commands[input]

		if !ok {
			fmt.Println("🤖 invalid command")
			continue
		}

		if command.name == "exit" || command.name == "quit" {
			command.callback()
			break
		}

		command.callback()
	}
}
