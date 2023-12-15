package main

import (
	"fmt"

	"github.com/seajoshc/gokedex/internal/pokeapi"
)

func commandMap() {
	pokeapi := pokeapi.NewClient()
	res, err := pokeapi.GetLocationAreas()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
