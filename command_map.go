package main

import "github.com/seajoshc/gokedex/internal/pokeapi"

func commandMap() error {
	pokeapi.GetLocationAreas()
	return nil
}
