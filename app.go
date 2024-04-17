package main

import (
	"net/http"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

var port = flag.String("addr", "8080", "http port")

func main() {

	flag.Parse()

	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "This is the api")
	})

	route.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})

	route.Run(fmt.Printf(":%s", *port))
}
