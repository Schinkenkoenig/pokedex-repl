package commands

import (
	"fmt"
)

func CommandHelp(c *Config, param string) error {
	helpText := `
Welcome to the Pokedex!
Usage:

    `
	fmt.Println(helpText)

	for _, cmd := range GetCliOptions() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}

	fmt.Println()
	return nil
}
