package services
import (
	// "errors"
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