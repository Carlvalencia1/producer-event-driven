package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "¡Hola, Gin!"})
	})

	r.Run(":8080") // Inicia el servidor en el puerto 8080
}
