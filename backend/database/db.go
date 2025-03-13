package database

import (
	"log"
	"os"
	"todo-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func ConnectDB(){
	dsn:=os.Getenv("DATABASE_URL")
	var err error
	DB,err=gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Fatal("Failed to connect to the database")
	}
	
	// log.Println("Connected to the database")
	err=DB.AutoMigrate(&models.Todo{})
	if err!=nil{
		log.Fatal("Failed to migrate the database")
	}
	log.Println("Database Migrated")
}