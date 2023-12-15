package main

import (
	"fmt"

	"github.com/seajoshc/gokedex/internal/pokeapi"
)

func commandMap() {
	pokeapiClient := pokeapi.NewClient()
	res, err := pokeapiClient.GetLocationAreas()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
