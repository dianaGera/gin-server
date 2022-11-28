package main

import (
	// "fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	// "errors"
)

// models
type User struct {
	ID   		int    `json: "id"`
	Activity 	string `json: "name"`
	Age	  		int    `json: "age"`
}

// database
var boredList = []User{}

// views
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, boredList)
}

func getRandomData(c *gin.Context) {
	url := "https://www.boredapi.com/api/activity"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	if res.Status != "200 OK" {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		println(string(body))
		c.Data(http.StatusOK, "application/json", body)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	}
	
}

func addUser(c *gin.Context) {
	var newUser User

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
	router.GET("/users", getUsers)
	router.GET("/random-data", getRandomData)
	router.POST("/add-user", addUser)

	router.Run("localhost:8080")
}
