package commands

import (
	"errors"
	"fmt"
	"time"

	catch "github.com/Schinkenkoenig/pokedex-repl/internal/catch"
	pokeapi "github.com/Schinkenkoenig/pokedex-repl/internal/pokeapi"
)

func CommandCatch(c *Config, param string) error {
	if c == nil {
		return errors.New("nil pointer on config")
	}

	if len(param) == 0 {
		return errors.New("need an pokemon name param")
	}

	uri := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-species/%s", param)
	pokeUri := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", param)
	go pokeapi.GET[pokeapi.PokemonResponse](c.Api, pokeUri)
	response, err := pokeapi.GET[pokeapi.PokemonSpeciesResponse](c.Api, uri)
	if err != nil {
		return err
	}

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
		c.Pokedex[response.Name] = *response
		fmt.Printf("You catched %s!\n", response.Name)
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
