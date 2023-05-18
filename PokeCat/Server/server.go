package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	world "github.com/simple-pokemon/go/PokeCat/World"
)

type KeyPressed struct {
	Key string
}

var number int

func StartServer() {

	router := gin.New()
	router.Use(CORSMiddleware())
	router.GET("/api/test", test)
	router.POST("/api/player/action", playerAction)
	router.GET("/api/world", getWorld)

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

func playerAction(c *gin.Context) {
	var key KeyPressed

	if err := c.ShouldBindJSON(&key); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playerAction := strings.ToLower(key.Key)
	if playerAction == "w" {
		number++
	} else if playerAction == "s" {
		number--
	}

	c.JSON(http.StatusOK, gin.H{"key": number})
}

// func getNumber(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"number": number})
// }

func getWorld(c *gin.Context) {
	world := world.NewWorld()
	c.JSON(http.StatusOK, gin.H{"world": world.WorldGrid})
}
