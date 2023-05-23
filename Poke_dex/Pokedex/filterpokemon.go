package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
var filtered_pokedex map[string][]PokemonStruct

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

var pokWithTypeSlice []PokemonType
var pokNamMap map[string]PokemonName

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
		err = json.Unmarshal(JSONContentInBytes, &pokNamMap)
	} else if nameList == "poketype" {
		err = json.Unmarshal(JSONContentInBytes, &pokWithTypeSlice)
	} else {
		err = json.Unmarshal(JSONContentInBytes, &pokedex)
	}
	if err != nil {
		panic(err)
	}
}

func main() {

	UnmarshalJSONFile("DataJSON/pokemonName.json", "pokename")
	// Read poketype JSON file
	UnmarshalJSONFile("DataJSON/PokemonCrawed.json", "poketype") // pokWithTypeSlice
	UnmarshalJSONFile("pokedex.json", "pokedex")
	fmt.Println("marshal done")

	// Filter and generate Pokedex
	filtered_pokedex = make(map[string][]PokemonStruct)
	pok_in_types_Counter := make(map[string]int)
	for typename, typePokList := range pokedex {
		for _, pokemonInPokedex := range typePokList {
			var type_list []string
			for _, pt := range pokWithTypeSlice { // pt {Form, PokemonID, PokemonName, Type}
				if pt.PokemonID == pokemonInPokedex.Id {
					type_list = pt.Type // get all types of input POKEMON
				} else if pt.PokemonID > pokemonInPokedex.Id {
					break
				}
			}

			if pok_in_types_Counter[typename] < 3 {
				isExisted := false
				for _, t := range type_list {
					for _, filtered_pok := range filtered_pokedex[t] {
						if filtered_pok.Id == pokemonInPokedex.Id {
							isExisted = true
							break
						}
					}
				}

				if isExisted == false {
					pok_in_types_Counter[typename]++
					filtered_pokedex[typename] = append(filtered_pokedex[typename], pokemonInPokedex)
				}
				for _, t := range type_list {
					if t == typename {
						continue
					}
					indexToRemove := -1
					for idx, p := range pokedex[t] {
						if p.Id == pokemonInPokedex.Id {
							indexToRemove = idx
							break
						}
					}
					if indexToRemove != -1 {
						pokedex[t] = append(pokedex[t][:indexToRemove], pokedex[t][indexToRemove+1:]...)
					}
				}
			} else {
				pokedex[typename] = filtered_pokedex[typename]
			}

		}
	}

	fmt.Println(len(filtered_pokedex))

	// save pokedex to JSON file
	file, err := os.Create("filtered_pokedex.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(filtered_pokedex); err != nil {
		panic(err)
	}
}
