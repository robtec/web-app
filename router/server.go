package router

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RunHTTPServer(ctx context.Context, port string) (err error) {

	route := gin.Default()

	tm := time.Now()

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

	return route.Run(fmt.Sprintf(":%s", port))
}
