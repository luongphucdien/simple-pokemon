package world

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sync"
	"time"

	entity "github.com/simple-pokemon/go/PokeCat/Entity"
)

const WORLD_SIZE = 100
const FREE_TILE_SYMBOL = "FT"
const PLAYER_SYMBOL = "#"
const POKEMON_SYMBOL = "&"

var WORLD World

type World struct {
	WorldGrid    [][]string
	PlayerList   map[string]entity.Player
	PokeDex      map[string]entity.Pokemon
	CurrPokemons []entity.Pokemon
	IDList       []string
	Mu           sync.RWMutex
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

	newWorld := World{
		WorldGrid:    make([][]string, WORLD_SIZE),
		PokeDex:      make(map[string]entity.Pokemon),
		PlayerList:   make(map[string]entity.Player),
		IDList:       make([]string, 0),
		CurrPokemons: make([]entity.Pokemon, 0),
	}

	newWorld.PokeDex = __loadPokeDex()
	newWorld.IDList = __extractPokeIDs(newWorld.PokeDex)
	newWorld.WorldGrid = __initWorldGrid()
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

func __fillWorldGrid(worldGrid [][]string) {
	for i := range [WORLD_SIZE]int{} {
		for j := range [WORLD_SIZE]int{} {
			worldGrid[i][j] = FREE_TILE_SYMBOL
		}
	}
}

func SpawnPokemons() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(WORLD.IDList), func(i, j int) { WORLD.IDList[i], WORLD.IDList[j] = WORLD.IDList[j], WORLD.IDList[i] })

	for _, ID := range WORLD.IDList {
		if len(WORLD.CurrPokemons) >= 50 {
			break
		} else {
			pokemon := WORLD.PokeDex[ID]
			pokemon.Coordinate.X = r.Intn(WORLD_SIZE-0) + 0
			pokemon.Coordinate.Y = r.Intn(WORLD_SIZE-0) + 0

		RECHECK:
			if WORLD.WorldGrid[pokemon.Coordinate.Y][pokemon.Coordinate.X] != FREE_TILE_SYMBOL {
				pokemon.Coordinate.X = r.Intn(WORLD_SIZE-0) + 0
				pokemon.Coordinate.Y = r.Intn(WORLD_SIZE-0) + 0
				goto RECHECK
			} else {
				goto CONTINUE
			}

		CONTINUE:
			WORLD.CurrPokemons = append(WORLD.CurrPokemons, pokemon)
			WORLD.WorldGrid[pokemon.Coordinate.Y][pokemon.Coordinate.X] = POKEMON_SYMBOL + ID
			pokemon.EffortValueYield[len(pokemon.EffortValueYield)-1] = 0.5 + r.Float32()*(1-0.5)
		}
	}
}

func __extractPokeIDs(pokeDex map[string]entity.Pokemon) []string {
	var IDs []string

	for key := range pokeDex {
		IDs = append(IDs, key)
	}

	return IDs
}

func __loadPokeDex() map[string]entity.Pokemon {
	var pokeDex map[string]entity.Pokemon

	pokeDexJSON, err := os.Open("./PokeCat/Entity/pokedexFinal.json")
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	defer pokeDexJSON.Close()

	byteValue, _ := ioutil.ReadAll(pokeDexJSON)

	json.Unmarshal([]byte(byteValue), &pokeDex)

	return pokeDex
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
	worldGrid[player.Coordinate.Y][player.Coordinate.X] = PLAYER_SYMBOL + "-" + player.Username
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
