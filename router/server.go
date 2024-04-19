package router

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gocolly/colly/v2"
)

func RunHTTPServer(ctx context.Context, port string) (err error) {

	router := gin.Default()

	tm := time.Now()

	cly := colly.NewCollector()

	// Find and visit all links
	cly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	cly.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
			"time":  fmt.Sprintf("%s", tm),
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

	router.GET("/scrape", func(c *gin.Context) {

		cly.Visit("http://go-colly.org/")

		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})

	return router.Run(fmt.Sprintf(":%s", port))
}
