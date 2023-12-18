package main

import (
	"fmt"
)

func commandMapb(c *config) error {
	fmt.Println("Location areas:")
	var pageUrl *string
	if c.previousLocationAreasPage != nil {
		pageUrl = c.previousLocationAreasPage
	}

	res, err := c.pokeapiClient.ListLocationAreas(pageUrl)
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
