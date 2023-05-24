package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

var damage map[string]AttackDmg

var playername []string

type AttackDmg struct {
	DmgOnType map[string]float32 `json:"damage_on_type"`
}

type Pokemon struct {
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
	Level            float32  `json:"level"`
	AccumulatedExp   float32  `json:"accumulated_exp"`
	Type             string   `json:"type"`
}

type Player struct {
	Id          int       `json:"id"`
	UserName    string    `json:"userName"`
	PassWord    string    `json:"passWord"`
	PokemonList []Pokemon `json:"pokemonList"`
}
type Battle struct {
	Player1          Player
	Player2          Player
	CurrentTurn      string
	currentAttackPok Pokemon
	defensePok       Pokemon
}

var PlayerList map[string]Player

func main() {
	PlayerList = make(map[string]Player)
	getDamageWhenAttackDB("damageOnAttack.json")
	// Start TCP server
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}

	fmt.Println("Server started, listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err)
			continue
		}

		go handleConnection(conn)
	}
}

func getDamageWhenAttackDB(filepath string) {
	JSONFile, err := os.Open(filepath)
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

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Read player data from client
	playerData, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println("Error reading player data: ", err)
		return
	}
	playerByte := []byte(playerData)
	var player Player
	playername = append(playername, player.UserName)
	err = json.Unmarshal(playerByte, &player)
	if err != nil {
		log.Println("Error decoding player data: ", err)
		return
	}
	fmt.Println(player)

	// Store player in PlayerList
	PlayerList[player.UserName] = player

	fmt.Println("Player", player.UserName, "joined the server")

	// Wait for another player to join
	for len(PlayerList) < 2 {
		time.Sleep(time.Second)
		fmt.Println("Waiting for 1 more player to start the game!!!!")
	}
	sendMessage(conn, "System is randomly pick Pokemon join the battle!!\n")

	// Select 3 pokemons for each player
	for _, p := range PlayerList {
		if len(p.PokemonList) < 3 {
			fmt.Println("Player ", p.UserName, " has less than 3 pokemons. Unable to start the battle.")
			return
		}
		if player, ok := PlayerList[p.UserName]; ok {
			player.PokemonList = selectPokemons(p.PokemonList, 3)
			var listPokSelected []string
			for _, v := range player.PokemonList {
				listPokSelected = append(listPokSelected, v.Name+", ")
			}

			fmt.Println("Player ", player.UserName, "chooses following pokemon ", listPokSelected)
		}
	}

	player1, pok1, player2, pok2 := getPlayerOrder(PlayerList[playername[0]], PlayerList[playername[1]])
	battle := Battle{
		Player1:          player1,
		Player2:          player2,
		CurrentTurn:      player1.UserName,
		currentAttackPok: pok1,
		defensePok:       pok2,
	}
	fmt.Println("Battle started between ", player1.UserName, " and ", player2.UserName)
	fmt.Println("Player ", player1.UserName, " will make the first move")
	// Send battle start message to players
	sendMessage(conn, "Battle started")
	for {
		// Get the current player's turn
		currentPlayer := getPlayerByUsername(battle.CurrentTurn)
		opponentPlayer := getOpponentPlayer(currentPlayer)

		battle.defensePok.CurrentHP = battle.defensePok.CurrentHP - getAttackDamage(battle.currentAttackPok, battle.defensePok)

		if battle.defensePok.CurrentHP <= 0 {
			Pok := getPokeForBattle(opponentPlayer.PokemonList, battle.defensePok.Id)
			if Pok.Id < 0 {
				calculateExperiencePoints(currentPlayer, opponentPlayer)

			}
			battle.defensePok = Pok
		} else {
			battle.CurrentTurn = opponentPlayer.UserName
		}
		levelingPokAfterBattle(currentPlayer, opponentPlayer)
		// if(currentAttackPok )

	}
}
func getPlayerByUsername(username string) Player {
	for _, p := range PlayerList {
		if p.UserName == username {
			return p
		}
	}
	return Player{}
}
func getOpponentPlayer(currentPlayer Player) Player {
	for _, p := range PlayerList {
		if p.UserName != currentPlayer.UserName {
			return p
		}
	}
	return Player{}
}
func getPlayerOrder(p1 Player, p2 Player) (Player, Pokemon, Player, Pokemon) {
	pokP1 := getPokeForBattle(p1.PokemonList, -1)
	pokP2 := getPokeForBattle(p2.PokemonList, -1)
	isP1MoveFirst := decideFirstMove(pokP1, pokP2)
	if isP1MoveFirst == true {
		return p1, pokP1, p2, pokP2
	}
	return p2, pokP2, p1, pokP1
}
func sendMessage(conn net.Conn, message string) {
	conn.Write([]byte(message))
}
func selectPokemons(pokemons []Pokemon, count int) []Pokemon {
	if count > len(pokemons) {
		count = len(pokemons)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(pokemons), func(i, j int) {
		pokemons[i], pokemons[j] = pokemons[j], pokemons[i]
	})

	return pokemons[:count]
}

// p1 take move first: true otherwise false
func decideFirstMove(pok1 Pokemon, pok2 Pokemon) bool {
	var isP1MoveFirst bool
	if pok1.Speed > pok2.Speed {
		isP1MoveFirst = true
	} else if pok1.Speed < pok2.Speed {
		isP1MoveFirst = false
	} else {
		rand.Seed(time.Now().UnixNano())
		check := rand.Intn(1)
		if check == 1 {
			isP1MoveFirst = true
		} else {
			isP1MoveFirst = false
		}
	}
	return isP1MoveFirst
}

func calculateExperiencePoints(player1 Player, player2 Player) {
	totalExp := float32(0)

	for _, p := range player2.PokemonList {
		totalExp += p.AccumulatedExp
	}

	expPerPokemon := totalExp / float32(len(player2.PokemonList))

	for i := range player1.PokemonList {
		player1.PokemonList[i].AccumulatedExp += expPerPokemon / 3
	}

	fmt.Println("Experience points calculated and updated for", player1.UserName)
}
func getAttackDamage(pokAtk Pokemon, PokDfs Pokemon) float32 {
	rand.Seed(time.Now().UnixNano())
	isAttackNormal := rand.Intn(2)
	if isAttackNormal == 1 {
		return pokAtk.Attack - PokDfs.Defense
	}
	return pokAtk.SpecialAttack*damage[strconv.Itoa(pokAtk.Id)].DmgOnType[PokDfs.Type] - PokDfs.SpecialDefense
}

func getPokeForBattle(PokList []Pokemon, ex int) Pokemon {
	var ranidx int
	arrIdx := make([]int, 0)
	for {
		rand.Seed(time.Now().UnixNano())
		ranidx := rand.Intn(2)
		if PokList[ranidx].CurrentHP <= 0 {
			arrIdx = append(arrIdx, ranidx)
		}
		if ex != ranidx && PokList[ranidx].CurrentHP > 0 {
			break
		}
		if len(arrIdx) == 3 {
			ranidx = -1
		}
	}
	if ranidx == -1 {
		return Pokemon{}
	}
	return PokList[ranidx]
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

func levelingPokAfterBattle(p1 Player, p2 Player) {
	for _, p := range p1.PokemonList {
		EV := float32(p.EffortValueYield[len(p.EffortValueYield)-1])
		old_acuExp := p.AccumulatedExp
		p.AccumulatedExp = p.AccumulatedExp * (1 + EV)
		p.Attack = p.Attack * (1 + EV)
		p.MaxHP = p.MaxHP * (1 + EV)
		p.CurrentHP = p.MaxHP
		p.Defense = p.Defense * (1 + EV)
		p.SpecialAttack = p.SpecialAttack * (1 + EV)
		p.SpecialDefense = p.SpecialDefense * (1 + EV)
		if p.AccumulatedExp >= (old_acuExp * 2) {
			p.Level = p.Level + 1
		}
	}
	for _, p := range p2.PokemonList {
		EV := float32(p.EffortValueYield[len(p.EffortValueYield)-1])
		old_acuExp := p.AccumulatedExp
		p.AccumulatedExp = p.AccumulatedExp * (1 + EV)
		p.Attack = p.Attack * (1 + EV)
		p.MaxHP = p.MaxHP * (1 + EV)
		p.CurrentHP = p.MaxHP
		p.Defense = p.Defense * (1 + EV)
		p.SpecialAttack = p.SpecialAttack * (1 + EV)
		p.SpecialDefense = p.SpecialDefense * (1 + EV)
		if p.AccumulatedExp >= (old_acuExp * 2) {
			p.Level = p.Level + 1
		}
	}
}
