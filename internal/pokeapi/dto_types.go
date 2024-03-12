package pokeapi

import "fmt"

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
