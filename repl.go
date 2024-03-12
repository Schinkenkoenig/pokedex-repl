package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cliOptions := getCliOptions()

	for {
		fmt.Fprint(os.Stdout, PROMPT)

		scanner.Scan()

		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if cmd, ok := cliOptions[commandName]; ok {
			cmd.callback()
		} else {
			fmt.Println("unknown cmd")
			continue
		}

	}
}

type cliCommand struct {
	callback    func() error
	name        string
	description string
}

func getCliOptions() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(s string) []string {
	output := strings.ToLower(s)
	words := strings.Fields(output)
	return words
}
