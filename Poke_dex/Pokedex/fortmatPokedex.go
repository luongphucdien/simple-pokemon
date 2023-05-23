package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type PokemonStruct struct {
	Id               int      `json:"id"`
	Img              string   `json:"img_link"`
	Name             string   `json:"name"`
	BaseExperience   int      `json:"base_experience"`
	EffortValueYield []int    `json:"effort_value_yield"`
	Form             []string `json:"form"`
	Attack           int      `json:"attack"`
	Defense          int      `json:"defense"`
	SpecialAttack    int      `json:"special_attack"`
	SpecialDefense   int      `json:"special_defense"`
	Speed            int      `json:"speed"`
	MaxHP            int      `json:"max_hp"`
	CurrentHP        int      `json:"current_hp"`
}

type Pokemon struct {
	PokemonStruct
	Type string `json:"type"`
}

var pokemon_grouped map[string][]PokemonStruct

var pokedex map[string]Pokemon

var id_list []string

func main() {
	JSONFile, err := os.Open("final_pokedex.json")
	if err != nil {
		panic(err)
	}
	defer JSONFile.Close()
	// Read whole  file
	JSONContentInBytes, err := ioutil.ReadAll(JSONFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(JSONContentInBytes, &pokemon_grouped)
	pokedex = make(map[string]Pokemon)
	for typepok, typeList := range pokemon_grouped {
		for _, pok := range typeList {
			pokedex[strconv.Itoa(pok.Id)] = Pokemon{pok, typepok}
			id_list = append(id_list, strconv.Itoa(pok.Id))
		}
	}

	// Write output JSON file

	pokemon_groupedBytes, err := json.Marshal(pokedex)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("pokedexFinal.json", pokemon_groupedBytes, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Output written to pokedexFinal.json")
	for _, item := range id_list {
		fmt.Print(" '" + item + "', ")
	}

}

id_list := []string{ '197',  '198',  '215',  '25',  '26',  '100',  '56',  '57',  '62',  '92',  '93',  '94',  '13',  '14',  '15',  '63',  '64',  '65',  '4',  '5',  '6',  '12',  '21',  '22',  '16',  '17',  '18',  '74',  '75',  '76',  '7',  '8',  '9',  '10',  
'11',  '46',  '147',  '148',  '149',  '1',  '2',  '3',  '87',  '91',  '124',  '35',  '36',  '39',  '27',  '28',  '31',  '81',  '82',  '205'}