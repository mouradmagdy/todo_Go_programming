package main

import (
	"log"
	"todo-go/config"
"todo-go/database"
	"github.com/gin-gonic/gin"
	"todo-go/routes"
)

func main(){
	// load configuration file
	config.LoadEnv()

	// connect to database
	database.ConnectDB()

	// init Gin Router
	router:=gin.Default()

	// setup routes
	// function starting with capital letter is exported
	// and can be accessed from other packages unlike function starting with small letter
	// which is private to the package
	routes.SetupRoutes(router)

	router.Run(":8080")
	log.Println("Server started on port 8080")
}