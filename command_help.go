package main

import "fmt"

func commandHelp() error {
	helpText := `
Welcome to the Pokedex!
Usage:

    `
	fmt.Println(helpText)

	for _, cmd := range getCliOptions() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
