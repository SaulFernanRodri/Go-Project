package app

import (
	"myproject/config"
	"myproject/controllers"
	"myproject/internal/database"
	"myproject/repositories"
	"myproject/routes"
	"myproject/services"

	"github.com/gin-gonic/gin"
)

func InitializeApp() *gin.Engine {
	config.LoadConfig()
	db := database.InitDB()
	database.Migrate(db)

	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	router := routes.SetupRouter(userController)

	return router
}
