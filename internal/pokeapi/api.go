package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func GetLocationAreas(uri string) (*LocationAreaResponse, error) {
	body, err := pokeapiCall(uri)
	if err != nil {
		return nil, err
	}

	resp, err := toLoactionArea(body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func pokeapiCall(uri string) ([]byte, error) {
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body.Close()

	if res.StatusCode > 299 {
		err := fmt.Errorf("response failed with status code: %d, body: %s", res.StatusCode, body)
		return nil, err
	}
	return body, nil
}

func toLoactionArea(b []byte) (*LocationAreaResponse, error) {
	response := LocationAreaResponse{}
	err := json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
