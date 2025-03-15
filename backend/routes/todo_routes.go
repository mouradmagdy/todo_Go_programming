package routes
import (
	"github.com/gin-gonic/gin"
"todo-go/controllers"
"todo-go/middleware"
)

func SetupRoutes(router *gin.Engine){
	routes:=router.Group("/todos")
	routes.Use(middleware.AuthMiddleware()) // protect all routes
	{
		routes.GET("/",controllers.GetTodos)
		routes.GET("",controllers.GetTodos)
		routes.POST("/",controllers.CreateTodo)
		routes.PUT("/:id",controllers.UpdateTodo)
		routes.DELETE("/:id",controllers.DeleteTodo)
	}
}