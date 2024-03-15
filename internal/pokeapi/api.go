package pokeapi

import (
	"encoding/json"

	"internal/cache"
)

type PokeApi struct {
	cache *cache.Cache
}

func GET[T any](p *PokeApi, uri string) (*T, error) {
	respBody, err := p.get(uri)
	if err != nil {
		return nil, err
	}

	var r T
	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
