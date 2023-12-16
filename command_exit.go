package main

import "fmt"

func commandExit(c *config) error {
	fmt.Println("👋 Shutting down the Gokedex")
	return nil
}
