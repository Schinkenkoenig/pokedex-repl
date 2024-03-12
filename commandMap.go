package main

import (
	"errors"
	"fmt"

	"internal/pokeapi"
)

func commandMap(c *Config) error {
	if c == nil {
		return errors.New("nil pointer on config")
	}
	fmt.Printf("%v\n", c)
	defer fmt.Printf("%v\n", c)

	uri := c.next

	response, err := pokeapi.GetLocationAreas(uri)
	if err != nil {
		return err
	}

	response.PrintAreas()

	c.updateConfig(response.Next, response.Previous)

	return nil
}

func commandMapb(c *Config) error {
	if c == nil {
		return errors.New("nil pointer on config")
	}

	fmt.Printf("%v\n", c)
	defer fmt.Printf("%v\n", c)

	uri := c.previous

	if len(uri) == 0 {
		return errors.New("on the first page, cannot go back")
	}

	response, err := pokeapi.GetLocationAreas(uri)
	if err != nil {
		return err
	}

	response.PrintAreas()

	c.updateConfig(response.Next, response.Previous)

	return nil
}

func (c *Config) updateConfig(next, prev string) {
	c.next = next
	c.previous = prev
}
