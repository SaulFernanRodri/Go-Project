package main

import (
	"log"
	"my-backend/repository"
	"my-backend/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Conexión a la base de datos
	dsn := "host=localhost user=gorm dbname=gorm password=mypassword port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migración automática
	db.AutoMigrate(&domain.User{})

	// Crear el repositorio y el servicio
	userRepo := repository.NewUserRepositoryGorm(db)
	userService := service.NewUserService(userRepo)

	// Crear el servidor HTTP
	router := gin.Default()

	// Ruta para obtener todos los usuarios
	router.GET("/users", func(c *gin.Context) {
		users, err := userService.GetAllUsers()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, users)
	})

	// Iniciar el servidor en el puerto 8080
	router.Run(":8080")
}
