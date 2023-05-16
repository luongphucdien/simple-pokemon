package world

import (
	"encoding/gob"
	"os"

	entity "github.com/simple-pokemon/go/PokeCat/Entity"
)

const WORLD_SIZE = 10
const FREE_TILE_SYMBOL = "FT"
const PLAYER_SYMBOL = "#"
const POKEMON_SYMBOL = "&"

type World struct {
	WorldGrid [][]string
	PokeList  []entity.Pokemon
	Player    []entity.Player
}

type Coordinate struct {
	X int
	Y int
}

// IMPORTANT: RUN THIS ONCE
func NewWorld() World {
	newWorld := World{
		WorldGrid: __initWorldGrid(),
		PokeList:  make([]entity.Pokemon, 1),
		Player:    make([]entity.Player, 1),
	}
	return newWorld
}

func __initWorldGrid() [][]string {
	worldGrid := make([][]string, WORLD_SIZE)
	rows := make([]string, WORLD_SIZE*WORLD_SIZE)
	for i := 0; i < WORLD_SIZE; i++ {
		worldGrid[i] = rows[i*WORLD_SIZE : (i+1)*WORLD_SIZE : (i+1)*WORLD_SIZE]
	}
	__fillWorldGrid(worldGrid)
	return worldGrid
}

func __fillWorldGrid(worldGrid [][]string) [][]string {
	for i := range [WORLD_SIZE]int{} {
		for j := range [WORLD_SIZE]int{} {
			worldGrid[i][j] = FREE_TILE_SYMBOL
		}
	}
	SaveWorld(worldGrid)
	return worldGrid
}

func StartWorld() World{
	newWorld := World {
		WorldGrid: RetrieveEmptyWorld(),
		PokeList: make([]entity.Pokemon, 1),
		Player: make([]entity.Player, 1),
	}

	return newWorld
}

func spawnEntities(worldGrid [][]string, pokemonList []entity.Pokemon) [][]string {
	for _, pokemon := range pokemonList {
		worldGrid[pokemon.Coordinate.Y][pokemon.Coordinate.X] = POKEMON_SYMBOL
	}
	return worldGrid
}

func SaveWorld(worldGrid [][]string) {
	file, _ := os.Create("./PokeCat/World/world-empty.gob")
	defer file.Close()
	gob.NewEncoder(file).Encode(worldGrid)
}

func RetrieveEmptyWorld() [][]string {
	var emptyWorldGrid [][]string

	file, _ := os.Open("./PokeCat/World/world-empty.gob")
	defer file.Close()
	gob.NewDecoder(file).Decode(&emptyWorldGrid)

	return emptyWorldGrid
}
