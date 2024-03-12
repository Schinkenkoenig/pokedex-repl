package pokeapi

import (
	"encoding/json"

	"internal/cache"
)

type PokeApi struct {
	cache *cache.Cache
}

func (p *PokeApi) GetLocationAreas(uri string) (*LocationAreaResponse, error) {
	respBody, err := p.get(uri)
	if err != nil {
		return nil, err
	}

	r := LocationAreaResponse{}
	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (p *PokeApi) GetPokemonInLocation(uri string) (*LocationAreaResponseDetail, error) {
	respBody, err := p.get(uri)
	if err != nil {
		return nil, err
	}

	r := LocationAreaResponseDetail{}
	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
