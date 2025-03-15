package routes
import (
	"todo-go/controllers"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine){
	auth:=router.Group("/auth")
	{
		auth.POST("/register",controllers.Register)
		auth.POST("/login",controllers.Login)
	}
}