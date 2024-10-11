package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger es un middleware que registra la información básica de la solicitud.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtiene el tiempo antes de que comience el procesamiento
		startTime := time.Now()

		// Procesa la solicitud
		c.Next()

		// Calcula el tiempo total de procesamiento
		duration := time.Since(startTime)

		// Registra la información relevante
		log.Printf("%s %s %d %s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
	}
}
