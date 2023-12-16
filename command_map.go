package main

import (
	"fmt"
)

func commandMap(c *config) error {
	fmt.Println("Location areas:")

	res, err := c.pokeapiClient.GetLocationAreas(c.nextLocationAreasPage)
	if err != nil {
		fmt.Println(err)
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	c.nextLocationAreasPage = res.Next
	c.previousLocationAreasPage = res.Previous
	return nil
}
