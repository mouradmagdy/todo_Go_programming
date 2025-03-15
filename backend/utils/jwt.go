package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret=[]byte(os.Getenv("JWT_SECRET"))
type Claims struct{
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint)(string,error){
	expirationTime:=time.Now().Add(24*time.Hour)
	claims:=&Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string)(*Claims,error){
	Claims:=&Claims{}
	token,err:=jwt.ParseWithClaims(tokenString,Claims,func(token *jwt.Token)(interface{},error){
		return jwtSecret,nil
	})
	if err!=nil||!token.Valid{
		return nil,err
	}
	return Claims,nil
}