package controllers

import (
	"net/http"
	"todo-go/database"
	"todo-go/models"
	"todo-go/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c*gin.Context){
	var user models.User
	if err:=c.ShouldBindJSON(&user);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid input"})
		return
	}
	// hash pass before saving
	hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Error hashing password"})
		return
	}
	user.Password=string(hashedPassword)
	if err:=database.DB.Create(&user).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"User already exists"})
		return
	}
	c.JSON(http.StatusCreated,gin.H{"message":"User created successfully"})
}

func Login(c *gin.Context){
	var input models.User
	var user models.User

	if err:=c.ShouldBindBodyWithJSON(&input);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid input"})
		return
	}
	// find user by email
	if err:=database.DB.Where("email =?",input.Email).First(&user).Error;err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid credentials"})
		return
	}
	// compare password
	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(input.Password));err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid credentials"})
		return
	}
	// generate token
	token,err:=utils.GenerateToken(user.ID)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Error generating token"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"token":token})
}