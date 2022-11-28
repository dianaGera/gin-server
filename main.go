package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	// "errors"
)

// models
type BoredList struct {
	ID            int     `json: "id"`
	Activity      string  `json: "activity"`
	Type          string  `json: "type"`
	Participants  int     `json: "participants"`
	Price         float64 `json: "price"`
	Link          string  `json: "link"`
	Key           string  `json: "key"`
	Accessibility float32 `json: "accessibility"`
}

// database
var boredList = []BoredList{}
var chachedBoredList BoredList

// views
func getMyBoredList(c *gin.Context) {
	fmt.Printf("Results: %v\n", chachedBoredList)
	lastTask := "You didn't serched for any tasks yet"
	res := []string{
		"Your Bored list":  boredList,
		"Last viewd tasks": lastTask,
	}
	// if
	c.IndentedJSON(http.StatusOK, res)
}

func getRandomData(c *gin.Context) {
	url := "https://www.boredapi.com/api/activity"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	// res, err := http.Get(url)
	// if err != nil {
	// 	panic(err.Error())
	// }

	if res.Status == "200 OK" {
		body, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			panic(err.Error())
		}

		json.Unmarshal(body, &chachedBoredList)

		fmt.Printf("Results: %v\n", chachedBoredList)

		c.Data(http.StatusOK, "application/json", body)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": res.Status})
	}

}

func addMyBoredList(c *gin.Context) {
	var newUser BoredList

	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	boredList = append(boredList, newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}

// server conf
func main() {

	router := gin.Default()

	// routers
	router.GET("/mybored-list", getMyBoredList)
	router.GET("/random-bored-thing", getRandomData)

	router.POST("/add-bored", addMyBoredList)
	router.POST("/remove-bored", addMyBoredList)

	router.Run("localhost:8080")
}
