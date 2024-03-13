package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *Config, param string) error {
	if c == nil {
		return errors.New("nil pointer on config")
	}

	if len(param) == 0 {
		return errors.New("need an pokemon name param")
	}

	if _, ok := c.pokedex[param]; !ok {
		return fmt.Errorf("%s is not in the dex.", param)
	}

	uri := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", param)
	response, err := c.api.GetPokemon(uri)
	if err != nil {
		return err
	}

	response.PrintPokemon()

	return nil
}
