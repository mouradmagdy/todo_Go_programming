package main

import (
	"log"
	"todo-go/config"
	"todo-go/database"
	"todo-go/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	// load configuration file
	config.LoadEnv()

	// connect to database
	database.ConnectDB()

	// init Gin Router
	router:=gin.Default()

	// cors
router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:5173"}, 
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    AllowCredentials: true,
}))

	// setup routes
	// function starting with capital letter is exported
	// and can be accessed from other packages unlike function starting with small letter
	// which is private to the package
	routes.SetupAuthRoutes(router) // auth routes
	routes.SetupRoutes(router) // protected todo routes

	router.Run(":8080")
	log.Println("Server started on port 8080")
}