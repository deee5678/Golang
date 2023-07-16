package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record video", Completed: false},
}

// router.GET("/todo") calls below function to get the details
func getTodos(context *gin.Context) { // recieves http requests context of type *gin.context
	context.IndentedJSON(http.StatusOK, todos) // converts the context from struct format([]todo) to jsone format
}

func main() {
	router := gin.Default()       // router is our server
	router.GET("/todo", getTodos) // create end point for get api

	router.Run("localhost:9090") // to run the server router, so our application should be running on port 9090
}
