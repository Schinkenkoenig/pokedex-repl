package pokeapi

import (
	"encoding/json"
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

func (a *LocationAreaResponse) PrintAreas() {
	for _, area := range a.Results {
		fmt.Println(area.Name)
	}
}

func (api *PokeApi) GetLocationAreas(uri string) (*LocationAreaResponse, error) {
	body, err := api.pokeapiCall(uri)
	if err != nil {
		return nil, err
	}

	resp, err := toLocationArea(body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func toLocationArea(b []byte) (*LocationAreaResponse, error) {
	response := LocationAreaResponse{}
	err := json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
