package routes

import (
	"myproject/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {

	router := gin.Default()

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}

	return router
}
