package services

import (
	// "errors"
	"errors"
	"todo-go/database"
	"todo-go/models"
)

func GetAllTodos()[]models.Todo{
	var todos []models.Todo
	database.DB.Find(&todos)
	return todos
}

func CreateTodo(todo models.Todo)models.Todo{
	database.DB.Create(&todo)
	return todo
}

func UpdateTodo(id string,updatedTodo models.Todo)(models.Todo,error){
	var todo models.Todo
	if err:=database.DB.First(&todo,id).Error;err!=nil{
		return todo,errors.New("todo not found")
	}
	todo.Title=updatedTodo.Title
	todo.Done=updatedTodo.Done
	database.DB.Save(&todo)
	return todo,nil
}

func DeleteTodo(id string)error{
	var todo models.Todo
	if err:=database.DB.First(&todo,id).Error;err!=nil{
	return errors.New("todo not found")	}
	database.DB.Delete(&todo)
	return nil
}