package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	// load configuration file

	// connect to database

	// init Gin Router
	router:=gin.Default()
	router.Run(":8080")
	log.Println("Server started on port 8080")
}