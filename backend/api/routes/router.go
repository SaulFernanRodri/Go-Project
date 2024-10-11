package routes

import (
	"myproject/api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(router *gin.Engine, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	router.GET("/users", userController.GetAllUsers)
	router.POST("/users", userController.CreateUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)
}
