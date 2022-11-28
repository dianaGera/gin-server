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
var chachedBoredList = []BoredList{}

// views
func getMyBoredList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"Your List":   boredList,
		"Last search": chachedBoredList,
	})
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
		var data BoredList
		json.Unmarshal(body, &data)
		if len(chachedBoredList) > 0 {
			chachedBoredList[0] = data
		} else {
			chachedBoredList = append(chachedBoredList, data)
		}

		fmt.Printf("Results: %v\n", chachedBoredList)

		c.Data(http.StatusOK, "application/json", body)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": res.Status})
	}

}

// func getBoredItem(c *gin.Context) {
	
// }

func addMyBoredList(c *gin.Context) {
	if len(chachedBoredList) > 0 {
		boredList = append(boredList, chachedBoredList[0])
		c.IndentedJSON(http.StatusCreated, chachedBoredList[0])
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Nothing to add"})
	}
}

// server conf
func main() {

	router := gin.Default()

	// routers
	router.GET("/mybored-list", getMyBoredList)
	router.GET("/random-bored-thing", getRandomData)
	// router.GET("/get-bored", getBoredItem)

	router.POST("/add-bored", addMyBoredList)
	router.POST("/remove-bored", addMyBoredList)

	router.Run("localhost:8080")
}
