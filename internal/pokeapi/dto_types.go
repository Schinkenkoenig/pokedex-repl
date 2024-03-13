package pokeapi

import (
	"fmt"
)

type LocationAreaResponse struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	Count int `json:"count"`
}

type LocationAreaResponseDetail struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
	Id int `json:"id"`
}

type PokemonResponse struct {
	Name  string `json:"name"`
	Stats []struct {
		Stat struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"type"`
	} `json:"types"`
	Id     int `json:"id"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}

type PokemonSpeciesResponse struct {
	Name      string `json:"name"`
	Id        int    `json:"id"`
	CatchRate int    `json:"capture_rate"`
}

func (p *PokemonResponse) PrintPokemon() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)

	fmt.Println("Stats:")

	for _, s := range p.Stats {
		fmt.Printf("  - %s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
}

func (a *LocationAreaResponse) PrintAreas() {
	for _, area := range a.Results {
		fmt.Println(area.Name)
	}
}

func (a *LocationAreaResponseDetail) PrintPokemans() {
	for _, p := range a.PokemonEncounters {
		fmt.Printf("  - %s\n", p.Pokemon.Name)
	}
}
