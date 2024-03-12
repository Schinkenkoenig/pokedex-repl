module github.com/Schinkenkoenig/pokedex-repl

go 1.21.2

require internal/pokeapi v1.0.0

require internal/cache v1.0.0

replace internal/cache => ./internal/cache

replace internal/pokeapi => ./internal/pokeapi
