package router

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gocolly/colly/v2"
)

func RunHTTPServer(ctx context.Context, port string) (err error) {

	router := gin.Default()

	tm := time.Now()

	var msg = "all good"

	cly := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"),
		colly.AllowURLRevisit(),
	)

	cly.OnHTML("footer", func(e *colly.HTMLElement) {
		msg = strings.Replace(e.Text, "\n", "", -1)
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

		urlQ := c.Query("url")

		err := cly.Visit(urlQ)

		if err != nil {
			msg = err.Error()
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": msg,
			"time":  fmt.Sprintf("%s", tm),
		})
	})

	return router.Run(fmt.Sprintf(":%s", port))
}
