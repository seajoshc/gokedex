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

		if input == "" {
			commands["help"].callback()
			continue
		}

		command, ok := commands[input]

		if !ok {
			fmt.Println("🤖 invalid command")
			commands["help"].callback()
			continue
		}

		if input == "exit" || input == "quit" {
			command.callback()
			break
		}

		command.callback()
	}
}
