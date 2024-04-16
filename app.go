package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "This is the api, part 2")
	})

	route.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})

	route.Run(":8080")
}
