package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var todos = []todo{
	{ID: "1", Title: "First Todo", Description: "First Todo Description"},
	{ID: "2", Title: "Second Todo", Description: "Second Todo Description"},
	{ID: "3", Title: "Third Todo", Description: "Third Todo Description"},
}

func getAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func getTodo(c *gin.Context) {
	id := c.Param("id")

	for _, a := range todos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Todo not found"})
}

func postTodo(c *gin.Context) {
	var newTodo todo

	if error := c.BindJSON(&newTodo); error != nil {
		return
	}

	todos = append(todos, newTodo)

	c.IndentedJSON(http.StatusCreated, newTodo)
}

func RemoveIndex(s []todo, index int) []todo {
	return append(s[:index], s[index+1:]...)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	for i, todo := range todos {
		if todo.ID == id {
			todos = RemoveIndex(todos, i)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Todo not found"})
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")

	var updateData todo

	if error := c.BindJSON(&updateData); error != nil {
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = updateData
			c.IndentedJSON(http.StatusOK, todos[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Todo not found"})
}

func main() {
	r := gin.Default()

	r.GET("/todos", getAll)
	r.GET("/todos/:id", getTodo)
	r.POST("/todos/", postTodo)
	r.DELETE("/todos/:id", deleteTodo)
	r.PUT("/todos/:id", updateTodo)

	r.Run()
}
