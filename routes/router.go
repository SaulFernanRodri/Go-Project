package routes

import (
	"myproject/controllers"
	"myproject/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {

	router := gin.Default()

	router.Use(middlewares.BasicAuth())

	router.LoadHTMLGlob("templates/*")

	router.POST("/login", userController.Login)

	userRoutes := router.Group("/users", middlewares.AuthJWT())
	{
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}

	viewRoutes := router.Group("/s")
	{
		viewRoutes.GET("/users", userController.ShowAllUsers)
	}

	return router
}
