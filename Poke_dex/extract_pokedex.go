package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func AddPokemon(pok PokemonStruct) {
	var type_list []string
	for _, pt := range poketype { // pt {Form, PokemonID, PokemonName, Type}
		if pt.PokemonID == pok.Id {
			type_list = pt.Type // get all types of input POKEMON
		} else if pt.PokemonID > pok.Id {
			break
		}
	}
	for _, type_of_pok := range type_list {
		_, isExisted := pokedex[type_of_pok] // type already existed in pokedex
		if isExisted == false {
			pokedex[type_of_pok] = []PokemonStruct{}
		}
		fnd := false
		for _, poke_In_dex := range pokedex[type_of_pok] {
			if poke_In_dex.Id == pok.Id {
				fnd = true
				break
			}
		}
		if fnd == false {
			pokedex[type_of_pok] = append(pokedex[type_of_pok], pok)
		}
	}
}

func getHTMLDocument() *goquery.Document {
	resp, err := http.Get("http://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_effort_value_yield")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	return doc
}
func crawlPokedex() error {
	doc := getHTMLDocument()
	pokedex = make(map[string][]PokemonStruct)
	tables := doc.Find("table")
	targetTable := tables.Eq(1) // Find the target table

	// loop through each table row and extract Pokemon information
	targetTable.Find("tr").Each(func(rowIndex int, rowHtml *goquery.Selection) {

		columns := rowHtml.Find("td")

		if columns.Length() == 11 {
			// 	extract information from columns
			id, _ := strconv.Atoi(strings.TrimSpace(columns.Eq(0).Find("b").First().Text()))
			name := strings.TrimSpace(columns.Eq(2).Find("a").First().Text())
			img_tag := columns.Eq(1).Find("img")
			src, _ := img_tag.Attr("src")
			baseExp, _ := strconv.Atoi(strings.TrimSpace(columns.Eq(3).Text()))
			effortValues := make([]int, 7)
			for j := 0; j < 7; j++ {
				// HP, Attack, Defense, Sp.Attack, Sp.Defense, Speech and TotalEV
				value, _ := strconv.Atoi(strings.TrimSpace(columns.Eq(j + 4).Text()))
				effortValues[j] = value
			}

			// create Pokemon struct and add to pokedex
			pokemon := PokemonStruct{
				Id:               id,
				Img:              src,
				Name:             name,
				BaseExperience:   baseExp,
				EffortValueYield: effortValues,
				Form:             findAllForm(id),
				Attack:           1,
				Defense:          1,
				SpecialAttack:    1,
				SpecialDefense:   1,
				Speed:            1,
				MaxHP:            1,
				CurrentHP:        1,
				Level:            1,
				AccumulatedExp:   0,
			}
			AddPokemon(pokemon)
		}
	})

	// save pokedex to JSON file
	file, err := os.Create("pokedex.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(pokedex); err != nil {
		return err
	}

	return nil
}
func main() {
	groupPokemonsByType()
	crawlPokedex()

}

void setup() {
	for(int i = 13; i >= 8; i--){
		pinMode(i, OUTPUT); 
	}
	pinMode(7, INPUT);
	Serial.begin(9600);
  }
   
  // the loop function runs over and over again forever
  void loop() {
	uint32_t startTime = millis();
	uint32_t endTime = 10000;
	uint32_t period = 10000L;
   
	if (digitalRead(7) == HIGH) {
	  for (;millis() - startTime < period;) {
		digitalWrite(8, HIGH);   
		digitalWrite(9, HIGH);
		digitalWrite(10, HIGH);
		delay(100);
		digitalWrite(8, LOW);
		digitalWrite(9, LOW);
		digitalWrite(10, LOW);
		delay(100);
   
		digitalWrite(11, HIGH);   
		digitalWrite(12, HIGH);
		digitalWrite(13, HIGH);
		delay(100);
		digitalWrite(11, LOW);
		digitalWrite(12, LOW);
		digitalWrite(13, LOW);
		delay(100);
	  }
   
	  delay(10000);
   
	  startTime = millis();
	  for (;millis() - startTime < period;) {
		for (int i=0; i<3; i++) { // S
		  for (int j=8; j<=13; j++) {
			digitalWrite(j, HIGH);
		  } 
		  delay(500);
		  for (int j=8; j<=13; j++) {
			digitalWrite(j, LOW);
		  } 
		  delay(500);
		}
   
		for (int i=0; i<3; i++) { // O
		  for (int j=8; j<=13; j++) {
			digitalWrite(j, HIGH);
		  } 
		  delay(250);
		  for (int j=8; j<=13; j++) {
			digitalWrite(j, LOW);
		  } 
		  delay(250);
		}
   
		for (int i=0; i<3; i++) { // S
		  for (int j=8; j<=13; j++) {
			digitalWrite(j, HIGH);
		  } 
		  delay(500);
		  for (int j=8; j<=13; j++) {
			digitalWrite(j, LOW);
		  } 
		  delay(500);
		}
	  }  
	}
  }