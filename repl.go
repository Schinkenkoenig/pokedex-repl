package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"internal/cache"
	"internal/pokeapi"
)

const PROMPT = "Pokedex > "

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cliOptions := getCliOptions()
	cache := cache.NewCache(time.Second * 5)
	api := pokeapi.NewApi(cache)

	c := Config{
		previous: "",
		next:     "https://pokeapi.co/api/v2/location-area/",
		api:      api,
		pokedex:  make(map[string]pokeapi.PokemonSpeciesResponse),
	}

	for {
		fmt.Fprint(os.Stdout, PROMPT)

		scanner.Scan()

		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		var param string
		if len(words) == 2 {
			param = words[1]
		}

		if cmd, ok := cliOptions[commandName]; ok {
			err := cmd.callback(&c, param)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("unknown cmd")
			continue
		}

	}
}

type cliCommand struct {
	callback    func(c *Config, param string) error
	name        string
	description string
}

type Config struct {
	api      *pokeapi.PokeApi
	pokedex  map[string]pokeapi.PokemonSpeciesResponse
	next     string
	previous string
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
		"map": {
			name:        "map",
			description: "Show the next 20 map location.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 map location.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore {area-name}",
			description: "Explores the given area. Maybe you will find a shiny.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon-name}",
			description: "Trying to catch a pokemon. good luck",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect {pokemon-name}",
			description: "inpsect a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show what is in your pokedex",
			callback: func(c *Config, param string) error {
				fmt.Println("You catched: ")
				for p := range c.pokedex {
					fmt.Printf("  - %s\n", p)
				}
				return nil
			},
		},
	}
}

func cleanInput(s string) []string {
	output := strings.ToLower(s)
	words := strings.Fields(output)
	return words
}
