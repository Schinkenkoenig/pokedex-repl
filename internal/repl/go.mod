module repl

require (
	internal/pokeapi v1.0.0
	internal/commands v1.0.0
	internal/cache v1.0.0
)

replace (
	internal/pokeapi => ../pokeapi
	internal/commands => ../commands
	internal/cache => ../cache
)

go 1.22.1
