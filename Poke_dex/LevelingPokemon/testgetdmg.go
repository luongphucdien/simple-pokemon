package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var damage map[string]AttackDmg
var typeList []string

type AttackDmg struct {
	DmgOnType map[string]float32 `json:"damage_on_type"`
}

func main() {
	typeList = []string{"Ice", "Electric", "Psychic", "Water", "Fairy", "Fighting", "Grass", "Ground", "Flying",
		"Steel", "Fire", "Bug", "Dark", "Dragon", "Ghost", "Normal", "Poison", "Rock"}
	damage = make(map[string]AttackDmg)
	for i := 1; i <= 649; i++ {
		link := fmt.Sprintf("https://pokedex.org/#/pokemon/%d", i)
		attack := AttackDmg{
			DmgOnType: make(map[string]float32),
		}
		for _, v := range typeList {
			attack.DmgOnType[v] = float32(1)
		}
		damage[strconv.Itoa(i)] = getDamageWhenAttacked(link, attack)

		fmt.Println("Done ", strconv.Itoa(i))
	}
	for i := 650; i <= 999; i++ {
		attack := AttackDmg{
			DmgOnType: make(map[string]float32),
		}
		for _, v := range typeList {
			attack.DmgOnType[v] = float32(1)
		}
		damage[strconv.Itoa(i)] = attack

		fmt.Println("Done ", strconv.Itoa(i))
	}

	// Write output JSON file

	damageonAttackBytes, err := json.Marshal(damage)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("damageOnAttack.json", damageonAttackBytes, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Output written to damageOnAttack.json")
}

func getDamageWhenAttacked(link string, attack AttackDmg) AttackDmg {
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Define the regular expression pattern
	pattern := `([a-zA-Z]+)(\d+(\.\d+)?x)`
	regex := regexp.MustCompile(pattern)

	doc.Find("div.when-attacked").Each(func(i int, s *goquery.Selection) {
		doc.Find("div.when-attacked-row").Each(func(i int, s *goquery.Selection) {
			matches := regex.FindAllStringSubmatch(s.Text(), -1)
			for _, match := range matches {
				typeValue := match[1]
				value := match[2][:len(match[2])-1]
				typeValue = strings.ToUpper(string(typeValue[0])) + typeValue[1:]
				dmg, _ := strconv.ParseFloat(value, 32)
				attack.DmgOnType[typeValue] = float32(dmg)
			}
		})
	})
	return attack
}
