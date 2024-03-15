package commands

import (
	"errors"

	"internal/pokeapi"
)

func CommandMap(c *Config, param string) error {
	if c == nil {
		return errors.New("nil pointer on config")
	}
	uri := c.Next

	resp, err := pokeapi.GET[pokeapi.LocationAreaResponse](c.Api, uri)
	if err != nil {
		return err
	}

	resp.PrintAreas()

	c.updateConfig(resp.Next, resp.Previous)

	return nil
}

func CommandMapb(c *Config, param string) error {
	if c == nil {
		return errors.New("nil pointer on config")
	}

	uri := c.Previous

	if len(uri) == 0 {
		return errors.New("on the first page, cannot go back")
	}

	response, err := pokeapi.GET[pokeapi.LocationAreaResponse](c.Api, uri)
	if err != nil {
		return err
	}

	response.PrintAreas()

	c.updateConfig(response.Next, response.Previous)

	return nil
}

func (c *Config) updateConfig(next, prev string) {
	c.Next = next
	c.Previous = prev
}
