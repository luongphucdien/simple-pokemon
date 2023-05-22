package server

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	entity "github.com/simple-pokemon/go/PokeCat/Entity"
	world "github.com/simple-pokemon/go/PokeCat/World"
)

var number int

func StartServer() {

	router := gin.New()
	router.Use(CORSMiddleware())
	router.GET("/api/test", test)
	router.POST("/api/player/action", playerAction)
	router.GET("/api/world", getWorld)
	router.POST("/api/player", checkPlayer)
	router.POST("/api/player/offline", removePlayer)
	// router.GET("/api/player", getPlayer)

	router.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func test(c *gin.Context) {
	data := "Test"
	c.IndentedJSON(http.StatusOK, data)
}

type PlayerAction struct {
	Action   string
	Username string
}

func playerAction(c *gin.Context) {
	var playerAction PlayerAction

	if err := c.ShouldBindJSON(&playerAction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	keyPressed := strings.ToLower(playerAction.Action)
	player := world.WORLD.PlayerList[playerAction.Username]

	world.WORLD.Mu.Lock()

	if keyPressed == entity.W {
		if player.Coordinate.Y-1 >= 0 {
			world.WORLD.WorldGrid[player.Coordinate.Y][player.Coordinate.X] = world.FREE_TILE_SYMBOL
			player.Coordinate.Y--
			world.WORLD.WorldGrid[player.Coordinate.Y][player.Coordinate.X] = world.PLAYER_SYMBOL
		}
	} else if keyPressed == entity.S {
		if player.Coordinate.Y+1 < world.WORLD_SIZE {
			world.WORLD.WorldGrid[player.Coordinate.Y][player.Coordinate.X] = world.FREE_TILE_SYMBOL
			player.Coordinate.Y++
			world.WORLD.WorldGrid[player.Coordinate.Y][player.Coordinate.X] = world.PLAYER_SYMBOL
		}
	} else if keyPressed == entity.A {
		if player.Coordinate.X-1 >= 0 {
			world.WORLD.WorldGrid[player.Coordinate.Y][player.Coordinate.X] = world.FREE_TILE_SYMBOL
			player.Coordinate.X--
			world.WORLD.WorldGrid[player.Coordinate.Y][player.Coordinate.X] = world.PLAYER_SYMBOL
		}
	} else if keyPressed == entity.D {
		if player.Coordinate.X+1 < world.WORLD_SIZE {
			world.WORLD.WorldGrid[player.Coordinate.Y][player.Coordinate.X] = world.FREE_TILE_SYMBOL
			player.Coordinate.X++
			world.WORLD.WorldGrid[player.Coordinate.Y][player.Coordinate.X] = world.PLAYER_SYMBOL
		}
	}
	world.WORLD.PlayerList[playerAction.Username] = player

	world.WORLD.Mu.Unlock()

	c.JSON(http.StatusOK, gin.H{"world-data": world.WORLD})
}

// func getNumber(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"number": number})
// }

func getWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"world": world.WORLD.WorldGrid, "world_data": world.WORLD})
}

func checkPlayer(c *gin.Context) {
	var player entity.Player

	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := os.Stat("./PokeCat/Server/players/" + player.Username + ".gob"); err != nil {
		playerFile, err := os.Create("./PokeCat/Server/players/" + player.Username + ".gob")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		defer playerFile.Close()

		player = world.GenerateRandomCoordinate(player)

		enc := gob.NewEncoder(playerFile)
		enc.Encode(&player)

		world.AddPlayer(player)
		c.JSON(http.StatusCreated, gin.H{"msg": "This is a new player. Returns player", "player_state": "PLAYER_NEW", "player": player})
	} else {
		playerFile, err := os.Open("./PokeCat/Server/players/" + player.Username + ".gob")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		defer playerFile.Close()

		dec := gob.NewDecoder(playerFile)
		var existedPlayer entity.Player
		dec.Decode(&existedPlayer)

		if player.Password == existedPlayer.Password {
			world.AddPlayer(existedPlayer)
			c.JSON(http.StatusOK, gin.H{"msg": "Player exists. Returns existed player", "player_state": "PLAYER_OLD", "player": existedPlayer})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Player exists. Password unmatched", "player_state": "PLAYER_UNAUTHORIZED"})
		}
	}

}

func removePlayer(c *gin.Context) {
	var player entity.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player = world.WORLD.PlayerList[player.Username]

	world.WORLD.Mu.Lock()
	world.WORLD.WorldGrid[player.Coordinate.Y][player.Coordinate.X] = world.FREE_TILE_SYMBOL
	delete(world.WORLD.PlayerList, player.Username)
	world.WORLD.Mu.Unlock()

	playerFile, err := os.Create("./PokeCat/Server/players/" + player.Username + ".gob")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer playerFile.Close()

	fmt.Println(player.Coordinate.X, player.Coordinate.Y)

	enc := gob.NewEncoder(playerFile)
	enc.Encode(&player)

	c.JSON(http.StatusOK, gin.H{"world-data": world.WORLD, "msg": "Player offline and removed"})
}

// func getPlayer(c *gin.Context) {
// 	player, err := os.Open("./Server/players" + player.Username + ".gob")
// }
