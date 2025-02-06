package controllers

import (
	"ejemplo/practica/src/Users/application"
	"ejemplo/practica/src/Users/infraestructure"
	"encoding/json"
	"github.com/gin-gonic/gin"
	
	"time"
)

// GetUserHandlerChunked - Long Polling para obtener usuarios
func GetUserHandler(c *gin.Context) {
	// Inicializar repositorio y caso de uso
	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewGetUser(repo)

	// Configurar encabezados para chunked
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	encoder := json.NewEncoder(c.Writer)

	// Enviar primer chunk indicando que la solicitud está en proceso
	encoder.Encode(gin.H{"status": "Obteniendo los usuarios..."})
	c.Writer.Flush()
	time.Sleep(1 * time.Second) // Simular latencia en la consulta a la base de datos

	// Ejecutar la obtención de usuarios
	users, err := useCase.Execute()
	if err != nil {
		encoder.Encode(gin.H{"error": "No se pudo obtener los usuarios"})
		c.Writer.Flush()
		return
	}

	// Enviar segundo chunk con los usuarios obtenidos
	encoder.Encode(users)
	c.Writer.Flush()
}
