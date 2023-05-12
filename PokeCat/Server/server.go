package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestComp struct {
	Test string
}

func StartServer() {
	router := gin.New()
	router.Use(CORSMiddleware())
	router.GET("/api/test", test)

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
