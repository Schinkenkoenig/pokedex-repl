package commands

func GetCliOptions() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    CommandExit,
		},
		"map": {
			Name:        "map",
			Description: "Show the next 20 map location.",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Show the previous 20 map location.",
			Callback:    CommandMapb,
		},
		"explore": {
			Name:        "explore {area-name}",
			Description: "Explores the given area. Maybe you will find a shiny.",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch {pokemon-name}",
			Description: "Trying to catch a pokemon. good luck",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect {pokemon-name}",
			Description: "inpsect a caught pokemon",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Show what is in your pokedex",
			Callback:    CommandPokedex,
		},
	}
}
