package middleware

import (
	"net/http"
	"todo-go/utils"
	"strings"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader:=c.GetHeader("Authorization")
		if authHeader==""{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"No token provided"})
			c.Abort()
			return
		}
		tokenString:=strings.Split(authHeader," ")[1]
		claims,err:=utils.ValidateToken(tokenString)
		if err!=nil{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid token"})
			c.Abort()
			return
		}
		c.Set("userId",claims.UserID)
		c.Next()
	}
}