package world

import (
	"math/rand"
	"sync"
	"time"

	entity "github.com/simple-pokemon/go/PokeCat/Entity"
)

const WORLD_SIZE = 10
const FREE_TILE_SYMBOL = "FT"
const PLAYER_SYMBOL = "#"
const POKEMON_SYMBOL = "&"

var WORLD World
var PLAYERS_ONLINE []entity.Player
var POKEMON_LIST []entity.Pokemon

type World struct {
	WorldGrid  [][]string
	PokeList   []entity.Pokemon
	PlayerList map[string]entity.Player
	Mu         sync.RWMutex
}

type Coordinate struct {
	X int
	Y int
}

type FreeTile struct {
	Pokemon entity.Pokemon
	Players []entity.Player
}

func NewWorld() World {
	PLAYERS_ONLINE = make([]entity.Player, 0)
	POKEMON_LIST = make([]entity.Pokemon, 0)

	pokeListTest := []entity.Pokemon{
		{Name: "test1", Entity: entity.Entity{Coordinate: entity.Coordinate{X: 5, Y: 5}}},
		{Name: "test2", Entity: entity.Entity{Coordinate: entity.Coordinate{X: 6, Y: 4}}},
		{Name: "test3", Entity: entity.Entity{Coordinate: entity.Coordinate{X: 7, Y: 3}}},
		{Name: "test4", Entity: entity.Entity{Coordinate: entity.Coordinate{X: 8, Y: 2}}},
		{Name: "test5", Entity: entity.Entity{Coordinate: entity.Coordinate{X: 2, Y: 8}}},
	}

	newWorld := World{
		WorldGrid:  make([][]string, WORLD_SIZE),
		PokeList:   pokeListTest,
		PlayerList: make(map[string]entity.Player),
	}

	newWorld.WorldGrid = __initWorldGrid(pokeListTest)
	return newWorld
}

func __initWorldGrid(pokeList []entity.Pokemon) [][]string {
	worldGrid := make([][]string, WORLD_SIZE)
	rows := make([]string, WORLD_SIZE*WORLD_SIZE)
	for i := 0; i < WORLD_SIZE; i++ {
		worldGrid[i] = rows[i*WORLD_SIZE : (i+1)*WORLD_SIZE : (i+1)*WORLD_SIZE]
	}
	__fillWorldGrid(worldGrid)
	__spawnEntities(worldGrid, pokeList)
	return worldGrid
}

func __fillWorldGrid(worldGrid [][]string) {
	for i := range [WORLD_SIZE]int{} {
		for j := range [WORLD_SIZE]int{} {
			worldGrid[i][j] = FREE_TILE_SYMBOL
		}
	}
}

func __spawnEntities(worldGrid [][]string, pokemonList []entity.Pokemon) {
	for _, pokemon := range pokemonList {
		worldGrid[pokemon.Coordinate.Y][pokemon.Coordinate.X] = POKEMON_SYMBOL
	}
}

func AddPlayer(player entity.Player) {
	WORLD.PlayerList[player.Username] = player
	WORLD.Mu.Lock()
	__updateSinglePlayer(WORLD.WorldGrid, player)
	WORLD.Mu.Unlock()
}

func GenerateRandomCoordinate(player entity.Player) entity.Player {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	xCoord := r.Intn(WORLD_SIZE-0) + 0
	yCoord := r.Intn(WORLD_SIZE-0) + 0

RECHECK:
	if WORLD.WorldGrid[yCoord][xCoord] == "FT" {
		player.Coordinate.X = xCoord
		player.Coordinate.Y = yCoord
		return player
	} else {
		xCoord = r.Intn(WORLD_SIZE-0) + 0
		yCoord = r.Intn(WORLD_SIZE-0) + 0
		goto RECHECK
	}
}

func __updatePlayer(worldGrid [][]string, playerList map[string]entity.Player) {
	for _, player := range playerList {
		worldGrid[player.Coordinate.Y][player.Coordinate.X] = PLAYER_SYMBOL
	}
}

func __updateSinglePlayer(worldGrid [][]string, player entity.Player) {
	worldGrid[player.Coordinate.Y][player.Coordinate.X] = PLAYER_SYMBOL
}

func ExecutePlayerAction(worldGrid [][]string, player entity.Player, keyPressed string) {
	
}

// func StartWorld() World{
// 	newWorld := World {
// 		WorldGrid: RetrieveEmptyWorld(),
// 		PokeList: make([]entity.Pokemon, 1),
// 		Player: make([]entity.Player, 1),
// 	}

// 	return newWorld
// }

// func SaveWorld(worldGrid [][]string) {
// 	file, _ := os.Create("./PokeCat/World/world-empty.gob")
// 	defer file.Close()
// 	gob.NewEncoder(file).Encode(worldGrid)
// }

// func RetrieveEmptyWorld() [][]string {
// 	var emptyWorldGrid [][]string

// 	file, _ := os.Open("./PokeCat/World/world-empty.gob")
// 	defer file.Close()
// 	gob.NewDecoder(file).Decode(&emptyWorldGrid)

// 	return emptyWorldGrid
// }
