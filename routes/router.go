package routes

import (
	"myproject/controllers"
	//"myproject/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	//router.Use(middlewares.BasicAuth())

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}

	viewRoutes := router.Group("/view")
	{
		viewRoutes.GET("/", userController.ShowAllUsers)
	}

	return router
}
