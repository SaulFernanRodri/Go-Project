package routes

import (
	"myproject/controllers"
	"myproject/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.Use(middlewares.BasicAuth())

	router.POST("/login", userController.Login)
	router.POST("/register", userController.CreateUser)

	userRoutes := router.Group("/users", middlewares.AuthJWT())
	{
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}

	viewRoutes := router.Group("/view")
	{
		viewRoutes.GET("/users", userController.ShowAllUsers)
	}

	return router
}
