package routes
import (
	"github.com/gin-gonic/gin"
"todo-go/controllers"
)

func SetupRoutes(router *gin.Engine){
	routes:=router.Group("/todos")
	{
		routes.GET("/",controllers.GetTodos)
		routes.GET("",controllers.GetTodos)
		routes.POST("/",controllers.CreateTodo)
		// routes.PUT("/:id",controllers.UpdateTodo)
		// routes.DELETE("/:id",controllers.DeleteTodo)
	}
}