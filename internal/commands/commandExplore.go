package commands

import (
	"errors"
	"fmt"

	pokeapi "github.com/Schinkenkoenig/pokedex-repl/internal/pokeapi"
)

func CommandExplore(c *Config, param string) error {
	if c == nil {
		return errors.New("nil pointer on config")
	}

	if len(param) == 0 {
		return errors.New("need an area param")
	}

	uri := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", param)
	response, err := pokeapi.GET[pokeapi.LocationAreaResponseDetail](c.Api, uri)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", param)
	fmt.Println("Found pokemon: ")

	response.PrintPokemans()

	return nil
}
