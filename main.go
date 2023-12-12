package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("gokedex 🎮 ")

		scanner.Scan() // get next input
		input := scanner.Text()
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Quitting Gokedex 👋")
			break
		}

		if input == "help" || input == "?" || input == "" {
			fmt.Println("Gokedex commands are: ")
			fmt.Println("help - shows the help menu you're reading :)")
			fmt.Println("exit - quits Gokedex")
			continue
		}

		fmt.Println("🤖 invalid command")
	}
}
