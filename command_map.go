package main

import (
	"fmt"

	"github.com/seajoshc/gokedex/internal/pokeapi"
)

func commandMap() {
	fmt.Println("Location areas:")
	pokeapiClient := pokeapi.NewClient()
	res, err := pokeapiClient.GetLocationAreas()
	if err != nil {
		fmt.Println(err)
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	pokeapiClient.NextPage = res.Next
	pokeapiClient.PreviousPage = res.Previous
}
