package commands

import (
	"errors"
	"fmt"

	"internal/pokeapi"
)

func CommandInspect(c *Config, param string) error {
	if c == nil {
		return errors.New("nil pointer on config")
	}

	if len(param) == 0 {
		return errors.New("need an pokemon name param")
	}

	if _, ok := c.Pokedex[param]; !ok {
		return fmt.Errorf("%s is not in the dex", param)
	}

	uri := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", param)
	response, err := pokeapi.GET[pokeapi.PokemonResponse](c.Api, uri)
	if err != nil {
		return err
	}

	response.PrintPokemon()

	return nil
}
