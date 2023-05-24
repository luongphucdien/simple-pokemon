package main

import (
	"strings"
	"time"
	server "github.com/simple-pokemon/go/PokeCat/Server"
	world "github.com/simple-pokemon/go/PokeCat/World"
)
func main() {
	world.WORLD = world.NewWorld()
	world.SpawnPokemons()
	go func ()  {
		ticker := time.NewTicker(10*time.Second)
		for range ticker.C {
			for i, row := range world.WORLD.WorldGrid {
				for j, tile := range row {
					if strings.Contains(tile, "&") {
						world.WORLD.Mu.Lock()
						world.WORLD.WorldGrid[i][j] = world.FREE_TILE_SYMBOL
						world.WORLD.Mu.Unlock()
					}
				}
			}
			world.WORLD.Mu.Lock()
			world.WORLD.CurrPokemons = world.WORLD.CurrPokemons[:0]
			world.WORLD.Mu.Unlock()
			world.SpawnPokemons()
		}
	}()
	server.StartServer()
}