package commands

import "fmt"

func CommandPokedex(c *Config, param string) error {
	fmt.Println("You catched: ")
	for p := range c.Pokedex {
		fmt.Printf("  - %s\n", p)
	}
	return nil
}
