package controllers

import (
	"net/http"
	"todo-go/models"
	"todo-go/services"

	"github.com/gin-gonic/gin"
)

func GetTodos(c*gin.Context){
	todos:=services.GetAllTodos()
	c.JSON(http.StatusOK,todos)
}

func CreateTodo(c*gin.Context){
	var todo models.Todo
	if err:=c.ShouldBindJSON(&todo);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	createdTodo:=services.CreateTodo(todo)
	c.JSON(http.StatusCreated,createdTodo)
}

