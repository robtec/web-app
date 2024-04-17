package main

import (
	"time"
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	var (
		port        string
	)

	flag.StringVar(&port, "port", "8080", "http port")

	flag.Parse()

	tm := time.Now()

	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "This is the api")
	})

	route.GET("/deployed", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("App deployed at %s", tm))
	})

	route.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})

	route.Run(fmt.Sprintf(":%s", *port))
}
