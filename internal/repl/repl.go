package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"internal/cache"
	"internal/commands"
	"internal/pokeapi"
)

const PROMPT = "Pokedex > "

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cliOptions := commands.GetCliOptions()
	cache := cache.NewCache(time.Second * 5)
	api := pokeapi.NewApi(cache)

	c := commands.Config{
		Previous: "",
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Api:      api,
		Pokedex:  make(map[string]pokeapi.PokemonSpeciesResponse),
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

func cleanInput(s string) []string {
	output := strings.ToLower(s)
	words := strings.Fields(output)
	return words
}
