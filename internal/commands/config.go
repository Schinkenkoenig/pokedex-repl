package commands

import "internal/pokeapi"

type Config struct {
	Api      *pokeapi.PokeApi
	Pokedex  map[string]pokeapi.PokemonSpeciesResponse
	Next     string
	Previous string
}

type CliCommand struct {
	Callback    func(c *Config, param string) error
	Name        string
	Description string
}
