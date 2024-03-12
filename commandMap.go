package main

import (
	"errors"
	"fmt"
)

func commandMap(c *Config, param string) error {
	if c == nil {
		return errors.New("nil pointer on config")
	}
	fmt.Printf("%v\n", c)
	defer fmt.Printf("%v\n", c)

	uri := c.next

	resp, err := c.api.GetLocationAreas(uri)
	if err != nil {
		return err
	}

	resp.PrintAreas()

	c.updateConfig(resp.Next, resp.Previous)

	return nil
}

func commandMapb(c *Config, param string) error {
	if c == nil {
		return errors.New("nil pointer on config")
	}

	fmt.Printf("%v\n", c)
	defer fmt.Printf("%v\n", c)

	uri := c.previous

	if len(uri) == 0 {
		return errors.New("on the first page, cannot go back")
	}

	response, err := c.api.GetLocationAreas(uri)
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
