package main

import (
	"fmt"
)

func commandMap(c *config) error {
	fmt.Println("Location areas:")

	res, err := c.pokeapiClient.GetLocationAreas()
	if err != nil {
		fmt.Println(err)
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	c.nextPage = res.Next
	c.previousPage = res.Previous
	return nil
}
