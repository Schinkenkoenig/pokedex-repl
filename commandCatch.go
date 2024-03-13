package main

import (
	"errors"
	"fmt"
	"time"

	"internal/catch"
)

func commandCatch(c *Config, param string) error {
	if c == nil {
		return errors.New("nil pointer on config")
	}

	if len(param) == 0 {
		return errors.New("need an pokemon name param")
	}

	uri := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-species/%s", param)
	pokeUri := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", param)
	go c.api.GetPokemon(pokeUri)
	response, err := c.api.GetPokemonSpecies(uri)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", response)

	catcher, err := catch.NewCatcher(catch.BOOTDEV)
	if err != nil {
		return err
	}

	res := catcher.CanCatch(response.CatchRate)
	w1, w2, w3 := res.Wobbles()

	printWobble(w1)
	printWobble(w2)
	printWobble(w3)

	if res.IsFullCatch() {
		fmt.Printf("You catched %s!\n", response.Name)
		c.pokedex[response.Name] = *response
	} else {
		fmt.Printf("You failed to catch %s\n", response.Name)
	}

	return nil
}

func printWobble(b bool) {
	if b {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("wobble...")
	}
}
