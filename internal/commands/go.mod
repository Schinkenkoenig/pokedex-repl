module commands

require (
	internal/pokeapi v1.0.0
	internal/catch v1.0.0
	internal/cache v1.0.0
)

replace (
	internal/pokeapi => ../pokeapi
	internal/catch => ../catch
	internal/cache => ../cache
)

go 1.22.1
