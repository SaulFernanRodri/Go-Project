package main

import (
	"myproject/api/routes"
	"myproject/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa la base de datos
	db := database.Init()

	// Crea una nueva instancia de Gin
	r := gin.Default()

	// Configura las rutas
	routes.Setup(r, db)

	// Inicia el servidor
	r.Run(":8080")
}
