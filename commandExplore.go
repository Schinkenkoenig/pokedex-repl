package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *Config, param string) error {
	if c == nil {
		return errors.New("nil pointer on config")
	}

	if len(param) == 0 {
		return errors.New("need a area param")
	}

	uri := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", param)
	response, err := c.api.GetPokemonInLocation(uri)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", param)
	fmt.Println("Found pokemon: ")

	response.PrintPokemans()

	return nil
}
