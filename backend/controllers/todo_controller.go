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

func UpdateTodo(c*gin.Context){
	id:=c.Param("id");
	var todo models.Todo
	if err:=c.ShouldBindJSON(&todo); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	updatedTodo,err:=services.UpdateTodo(id,todo)
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"Todo not found"})
		return
	}
	c.JSON(http.StatusOK,updatedTodo)
}

func DeleteTodo(c*gin.Context){
	id:=c.Param("id")
 	if err:=services.DeleteTodo(id);err!=nil{
	 c.JSON(http.StatusNotFound,gin.H{"error":"Todo not found"})
	 return
 }
 c.JSON(http.StatusOK,gin.H{"message":"Todo deleted"})
}
