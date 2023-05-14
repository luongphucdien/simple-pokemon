package main

import (
	server "github.com/simple-pokemon/go/PokeCat/Server"
	world "github.com/simple-pokemon/go/PokeCat/World"
)

func main() {	
	world.WORLD = world.NewWorld()
	world.SpawnPokemons()
	// world.SpawnTestPokemon()
	
	server.StartServer()
}