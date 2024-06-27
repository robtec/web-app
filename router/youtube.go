package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Kind  string  `json:"kind"`
	Items []Items `json:"items"`
}

type Items struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}

func Youtube(c *gin.Context) {

	var response Response

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if _, ok := c.GetQueryArray("key"); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"missing parameter": "key",
		})
		return
	}

	if _, ok := c.GetQueryArray("id"); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"missing parameter": "id",
		})
		return
	}

	q := req.URL.Query()

	q.Add("key", c.Query("key"))
	q.Add("id", c.Query("id"))
	q.Add("part", "statistics")

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)

	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &response)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	resultCode := resp.StatusCode

	if resultCode >= 300 || resultCode <= 199 {
		c.JSON(resultCode, response)
		return
	}

	if len(response.Items) > 0 {
		numberStr := response.Items[0].Stats.Subscribers

		number, _ := strconv.Atoi(numberStr)

		c.JSON(http.StatusOK, gin.H{
			"number": number,
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
