package router

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RunHTTPServer(ctx context.Context, port string) (err error) {

	router := gin.Default()

	tm := time.Now()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
			"time": fmt.Sprintf("%s", tm),
		})
	})

	router.GET("/deployed", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("App deployed at %s", tm))
	})

	router.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})

	return router.Run(fmt.Sprintf(":%s", port))
}
