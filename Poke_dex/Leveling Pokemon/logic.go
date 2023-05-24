package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var damage map[string]AttackDmg

type AttackDmg struct {
	DmgOnType map[string]float32 `json:"damage_on_type"`
}

type Pokemon struct {
	Id               int      `json:"id"`
	Type             string   `json:"type"`
	Img              string   `json:"img_link"`
	Name             string   `json:"name"`
	BaseExperience   int      `json:"base_experience"`
	EffortValueYield []int    `json:"effort_value_yield"`
	Form             []string `json:"form"`
	Attack           float32  `json:"attack"`
	Defense          float32  `json:"defense"`
	SpecialAttack    float32  `json:"special_attack"`
	SpecialDefense   float32  `json:"special_defense"`
	Speed            int      `json:"speed"`
	MaxHP            float32  `json:"max_hp"`
	CurrentHP        float32  `json:"current_hp"`
	Level            float32  `json:"level"`
	AccumulatedExp   float32  `json:"accumulated_exp"`
}
type Pokedex struct {
	Id               int      `json:"id"`
	Img              string   `json:"img_link"`
	Name             string   `json:"name"`
	BaseExperience   int      `json:"base_experience"`
	EffortValueYield []int    `json:"effort_value_yield"`
	Form             []string `json:"form"`
	Attack           float32  `json:"attack"`
	Defense          float32  `json:"defense"`
	SpecialAttack    float32  `json:"special_attack"`
	SpecialDefense   float32  `json:"special_defense"`
	Speed            int      `json:"speed"`
	MaxHP            float32  `json:"max_hp"`
	CurrentHP        float32  `json:"current_hp"`
}
type Player struct {
	Id          int       `json:"id"`
	UserName    string    `json:"userName"`
	PassWord    string    `json:"passWord"`
	PokemonList []Pokemon `json:"pokemonList"`
}

var PlayerList map[string]Player

func addPokemonToPlayerList(pokedex Pokedex, playerId string, typepok string) {
	if player, ok := PlayerList[playerId]; ok {
		// Initial a catched pokemon to a player with level 1 and accumulate = 0
		pok := Pokemon{pokedex.Id, typepok, pokedex.Img, pokedex.Name, pokedex.BaseExperience, pokedex.EffortValueYield, pokedex.Form,
			pokedex.Attack, pokedex.Defense, pokedex.SpecialAttack, pokedex.SpecialDefense, pokedex.Speed, pokedex.MaxHP, pokedex.CurrentHP, 1, 0}
		player.PokemonList = append(player.PokemonList, pok)
		PlayerList[playerId] = player
	}
}

func getPokeForBattle(idxPick int, PokList []Pokemon) Pokemon {
	var pok Pokemon
	for idx, v := range PokList {
		if idx == idxPick {
			return v
		}
	}
	return pok
}

// switch pok if another is kill

func switchOtherPok(PokList []Pokemon) {

}

// func levelingPokemon(playerId string, pokPickedList []Pokemon, expPoint int, accumulatedExp int, ev float32){
// 	if player, isExisted := PlayerList[playerId]; isExisted {
// 		for _, pok := range player.PokemonList{
// 			for _, pickedPok := range pokPickedList{
// 				if pok.Id == pic
// 			}
// 		}
// 	}
// }

// p1 take move first: true otherwise false
func decideFirstMove(pok1 Pokemon, pok2 Pokemon) bool {
	var isP1MoveFirst bool
	if pok1.Speed > pok2.Speed {
		isP1MoveFirst = true
	} else if pok1.Speed < pok2.Speed {
		isP1MoveFirst = false
	} else {
		rand.Seed(time.Now().UnixNano())
		check := rand.Intn(2)
		if check == 1 {
			isP1MoveFirst = true
		} else {
			isP1MoveFirst = false
		}
	}
	return isP1MoveFirst
}
func getAttackDamage(pokAtk Pokemon, PokDfs Pokemon) float32 {
	rand.Seed(time.Now().UnixNano())
	isAttackNormal := rand.Intn(2)
	if isAttackNormal == 1 {
		return pokAtk.Attack - PokDfs.Defense
	}
	return pokAtk.SpecialAttack*damage[strconv.Itoa(pokAtk.Id)].DmgOnType[PokDfs.Type] - PokDfs.SpecialDefense
}

func destroyToAllowInherit(pok Pokemon, pokeListPicked []Pokemon) bool {
	for _, pokPicked := range pokeListPicked {
		if pok.Id != pokPicked.Id {
			if pok.Type == pokPicked.Type {
				pokPicked.AccumulatedExp += pok.AccumulatedExp
				pok.AccumulatedExp = 0
				pok.CurrentHP = 0
				return true
			}
		}
	}
	return false
}

func pokeBat(pokeListP1 []Pokemon, pokeListP2 []Pokemon, IdPickP1 int, IdPickP2 int) {
	pokP1 := getPokeForBattle(IdPickP1, pokeListP1)
	pokP2 := getPokeForBattle(IdPickP2, pokeListP2)
	isP1MoveFirst := decideFirstMove(pokP1, pokP2)
	fmt.Println(isP1MoveFirst)
}

func main() {
	PlayerList = make(map[string]Player, 0)
	// fmt.Println(len(PlayerList))
	// Open JSON file
	JSONFile, err := os.Open("damageOnAttack.json")
	if err != nil {
		panic(err)
	}
	defer JSONFile.Close()
	// Read whole  file
	JSONContentInBytes, err := ioutil.ReadAll(JSONFile)
	err = json.Unmarshal(JSONContentInBytes, &damage)
	if err != nil {
		panic(err)
	}
}
