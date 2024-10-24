package routers

import (
	"myproject/controllers"
	//"myproject/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyProject API
// @version 1.0
// @description This is a sample server for MyProject API.
// @openapi:version 3.0.0
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.myproject.com/support
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8082
// @BasePath /api/v1
// @schemes http

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
