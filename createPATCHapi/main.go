package main

import (
	"errors"
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

// router.GET("/todo") endpoint calls below function to get the details
func getTodos(context *gin.Context) { // recieves http requests context of type *gin.context
	context.IndentedJSON(http.StatusOK, todos) // converts the context from struct format([]todo) to jsone format
}

// POST api's "/todos" end point, calls this functn
func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil { // context.BindJSON recieves the request thats in json format & binds the json with variable newTodo of type todo struct.(basically converting json to type todo. this might return an error if the json doesnt have the format of type struct check the declaration part. so we are 1st checking if it has any error, if yes we are just catching the error in err & returning it.)
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

// GET api to get specific id details
func getTodo(context *gin.Context) { // recieves http requests context of type *gin.context
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)

}

// PATCH api
func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)

}

// function to return specific id details
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()           // router is our server
	router.GET("/todo", getTodos)     // create end point for get api
	router.GET("/todos/:id", getTodo) // create end point for GET api by id
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo) // create end point for POST api
	router.Run("localhost:9090")   // to run the server router, so our application should be running on port 9090
}
