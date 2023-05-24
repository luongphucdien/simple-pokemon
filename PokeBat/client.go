package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type Player struct {
	Id          int       `json:Id`
	UserName    string    `json:"userName"`
	PassWord    string    `json:"passWord"`
	PokemonList []Pokemon `json:"pokemonList"`
}

type Pokemon struct {
	Id               int      `json:Id`
	Img              string   `json:Img`
	Name             string   `json:Name`
	BaseExperience   int      `json:BaseExperience`
	EffortValueYield []int    `json:EffortValueYield`
	Form             []string `json:Form`
	Attack           float32  `json:Attack`
	Defense          float32  `json:Defense`
	SpecialAttack    float32  `json:SpecialAttack`
	SpecialDefense   float32  `json:SpecialDefense`
	Speed            int      `json:Speed`
	MaxHP            float32  `json:MaxHP`
	CurrentHP        float32  `json:CurrentHP`
	Type             string   `json:Type`
	Level            float32  `json:"level"`
	AccumulatedExp   float32  `json:"accumulated_exp"`
}

func main() {
	player1 := Player{
		Id:       1,
		UserName: "Player1",
		PassWord: "password",
		PokemonList: []Pokemon{
			{
				Id:             13,
				Img:            "//archives.bulbagarden.net/media/upload/thumb/3/36/0013Weedle.png/70px-0013Weedle.png",
				Name:           "Weedle",
				BaseExperience: 39,
				EffortValueYield: []int{
					0, 0, 0, 0, 0, 1, 1,
				},
				Form:           []string{"Normal"},
				Attack:         35,
				Defense:        30,
				SpecialAttack:  20,
				SpecialDefense: 20,
				Speed:          50,
				MaxHP:          40,
				CurrentHP:      40,
				Type:           "Poison",
			},
			{
				Id:             14,
				Img:            "//archives.bulbagarden.net/media/upload/thumb/f/f3/0014Kakuna.png/70px-0014Kakuna.png",
				Name:           "Kakuna",
				BaseExperience: 72,
				EffortValueYield: []int{
					0, 0, 2, 0, 0, 0, 2,
				},
				Form:           []string{"Normal"},
				Attack:         25,
				Defense:        50,
				SpecialAttack:  25,
				SpecialDefense: 25,
				Speed:          35,
				MaxHP:          45,
				CurrentHP:      45,
				Type:           "Poison",
			},
			{
				Id:             147,
				Img:            "//archives.bulbagarden.net/media/upload/thumb/a/ae/0147Dratini.png/70px-0147Dratini.png",
				Name:           "Dratini",
				BaseExperience: 60,
				EffortValueYield: []int{
					0, 1, 0, 0, 0, 0, 1,
				},
				Form:           []string{"Normal"},
				Attack:         64,
				Defense:        45,
				SpecialAttack:  50,
				SpecialDefense: 50,
				Speed:          50,
				MaxHP:          41,
				CurrentHP:      41,
				Type:           "Dragon",
			},
		},
	}
	player2 := Player{
		Id:       2,
		UserName: "Player2",
		PassWord: "password",
		PokemonList: []Pokemon{
			{
				Id:             1,
				Img:            "//archives.bulbagarden.net/media/upload/thumb/f/fb/0001Bulbasaur.png/70px-0001Bulbasaur.png",
				Name:           "Bulbasaur",
				BaseExperience: 64,
				EffortValueYield: []int{
					0, 0, 0, 1, 0, 0, 1,
				},
				Form:           []string{"Normal"},
				Attack:         49,
				Defense:        49,
				SpecialAttack:  65,
				SpecialDefense: 65,
				Speed:          45,
				MaxHP:          45,
				CurrentHP:      45,
				Type:           "Grass",
			},
			{
				Id:             10,
				Img:            "//archives.bulbagarden.net/media/upload/thumb/5/5e/0010Caterpie.png/70px-0010Caterpie.png",
				Name:           "Caterpie",
				BaseExperience: 39,
				EffortValueYield: []int{
					1, 0, 0, 0, 0, 0, 1,
				},
				Form:           []string{"Normal"},
				Attack:         30,
				Defense:        35,
				SpecialAttack:  20,
				SpecialDefense: 20,
				Speed:          45,
				MaxHP:          45,
				CurrentHP:      45,
				Type:           "Bug",
			},
			{
				Id:             100,
				Img:            "//archives.bulbagarden.net/media/upload/thumb/5/55/0100Voltorb.png/70px-0100Voltorb.png",
				Name:           "Voltorb",
				BaseExperience: 66,
				EffortValueYield: []int{
					0, 0, 0, 0, 0, 1, 1,
				},
				Form:           []string{"Hisuian", "Normal"},
				Attack:         1,
				Defense:        1,
				SpecialAttack:  1,
				SpecialDefense: 1,
				Speed:          1,
				MaxHP:          1,
				CurrentHP:      1,
				Type:           "Electric",
			},
		},
	}
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error connecting to the server: ", err)
	}

	defer conn.Close()
	fmt.Println(player1)

	// Send player data to the server
	playerData, err := json.Marshal(player2)
	if err != nil {
		log.Fatal("Error encoding player data: ", err)
	}
	playerData = append(playerData, byte('\n'))
	// conn.Write(playerData)
	fmt.Fprintf(conn, string(playerData))

	// Wait for battle result from the server
	message, err := readMessage(conn)
	if err != nil {
		log.Fatal("Error reading battle result: ", err)
	}

	fmt.Println("System messages:", message)

	// Wait for battle result from the server
	message, err = readMessage(conn)
	if err != nil {
		log.Fatal("Error reading battle result: ", err)
	}

	fmt.Println("System messages:", message)
}

func readMessage(conn net.Conn) (string, error) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:n]), nil
}
