package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/exp/slices"
)

type PokemonName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PokemonType struct {
	Form        string   `json:"Form"`
	PokemonID   int      `json:"PokemonID"`
	PokemonName string   `json:"PokemonName"`
	Type        []string `json:"Type"`
}
type Pokemon struct {
	ID   int      `json:"id"`
	Name string   `json:"name"`
	Form []string `json:"Form"`
}

var pokedex map[string][]PokemonStruct

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
	Level            int      `json:"level"`
	AccumulatedExp   int      `json:"accumulated_exp"`
}

// type Pokemon map[string][]Pokemon
var poketype []PokemonType
var pokename map[string]PokemonName

// Group Pokémon by type
var pokemon_grouped map[string][]Pokemon

// check whether a pokemon existed in the type_list( contains pokemon with same type) or not
func isPokExistInTypeList(pok PokemonName, typePokemon string) bool {
	for _, p := range pokemon_grouped[typePokemon] {
		if p.ID == pok.ID {
			return true
		}
	}
	return false
}

// List all the forms of each Pokemon in the list
func findAllForm(ID int) []string {
	var form_list []string
	for _, poke := range poketype {
		if ID == poke.PokemonID {
			isContain := slices.Contains(form_list, poke.Form)
			if isContain == false {
				form_list = append(form_list, poke.Form)
			}
		} else if poke.PokemonID > ID {
			break
		}
	}
	return form_list
}

func UnmarshalJSONFile(file_path string, nameList string) {
	// Open JSON file
	JSONFile, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer JSONFile.Close()
	// Read whole  file
	JSONContentInBytes, err := ioutil.ReadAll(JSONFile)
	if err != nil {
		panic(err)
	}
	if nameList == "pokename" {
		err = json.Unmarshal(JSONContentInBytes, &pokename)
	} else {
		err = json.Unmarshal(JSONContentInBytes, &poketype)
	}
	if err != nil {
		panic(err)
	}
}
func groupPokemonsByType() {
	// Unmarshal pokename file into pokename that list all name of released pokemons with their ID
	UnmarshalJSONFile("DataJSON/pokemonName.json", "pokename")
	// Read poketype JSON file
	UnmarshalJSONFile("DataJSON/PokemonCrawed.json", "poketype")
	pokemon_grouped = make(map[string][]Pokemon)

	// Iterate each pokemon.
	idx := 0
	for _, pok_Name := range pokename { // ID, Name
		form_list := findAllForm(pok_Name.ID)
		if idx == pok_Name.ID {
			continue
		}
		// check each pokemon in file pokemon list with type and form
		for _, pok_Type_List := range poketype { // Form, PokemonID, PokemonName, Type
			if pok_Type_List.PokemonID == pok_Name.ID {
				for _, type_add := range pok_Type_List.Type {
					_, isTypeExisted := pokemon_grouped[type_add]
					pok := Pokemon{pok_Name.ID, pok_Name.Name, form_list}
					if isTypeExisted == false {
						pokemon_grouped[type_add] = []Pokemon{}
					}
					if isPokExistInTypeList(pok_Name, type_add) == false {
						// the pokemon does not exist in the type_list
						pokemon_grouped[type_add] = append(pokemon_grouped[type_add], pok)
					}
				}
			}
		}
		idx = idx + 1
	}

	// Write output JSON file
	pokemon_groupedBytes, err := json.Marshal(pokemon_grouped)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("pokemon.json", pokemon_groupedBytes, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Output written to pokemon.json")
}
