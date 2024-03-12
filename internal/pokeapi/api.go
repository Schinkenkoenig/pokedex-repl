package pokeapi

import (
	"fmt"
	"io"
	"net/http"

	"internal/cache"
)

type PokeApi struct {
	cache *cache.Cache
}

func NewApi(cache *cache.Cache) *PokeApi {
	api := PokeApi{cache: cache}
	return &api
}

func (p *PokeApi) pokeapiCall(uri string) ([]byte, error) {
	cacheBody, ok := p.cache.Get(uri)

	if ok {
		return cacheBody, nil
	}

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

	p.cache.Add(uri, body)

	return body, nil
}
