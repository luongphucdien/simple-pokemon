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

func main() {
	// Read pokename JSON file
	pokenameFile, err := os.Open("DataJSON/pokemonName.json")
	if err != nil {
		panic(err)
	}
	defer pokenameFile.Close()

	pokenameBytes, err := ioutil.ReadAll(pokenameFile)
	if err != nil {
		panic(err)
	}

	pokename := make(map[string]PokemonName)
	err = json.Unmarshal(pokenameBytes, &pokename)
	if err != nil {
		panic(err)
	}

	// Read poketype JSON file
	poketypeFile, err := os.Open("DataJSON/PokemonCrawed.json")
	if err != nil {
		panic(err)
	}
	defer poketypeFile.Close()

	poketypeBytes, err := ioutil.ReadAll(poketypeFile)
	if err != nil {
		panic(err)
	}

	poketype := make([]PokemonType, 0)
	err = json.Unmarshal(poketypeBytes, &poketype)
	if err != nil {
		panic(err)
	}
	// fmt.Println(pokename)

	// Group Pokémon by type
	grouped := make(map[string][]PokemonName)
	// fmt.Println(grouped)

	for _, p := range pokename {
		groupedType := make([]PokemonName, 0)
		for _, t := range poketype {
			if t.PokemonName == p.Name {
				fnd := false
				for _, typeCheck := range groupedType {
					if typeCheck.ID == p.ID {
						fnd = true
						break
					}
				}
				if fnd == false{

					groupedType = append(groupedType, p)
					
					grouped[t.Type[0]] = append(grouped[t.Type[0]], groupedType...) // append 
				}
			}
		}
	}
	

	// Write output JSON file
	groupedBytes, err := json.Marshal(grouped)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("pokemon.json", groupedBytes, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Output written to pokemon.json")
}