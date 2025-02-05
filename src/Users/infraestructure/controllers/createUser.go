package controllers

import (
	"ejemplo/practica/src/Users/application"
	"ejemplo/practica/src/Users/domain"
	"ejemplo/practica/src/Users/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CreateUsersHandler maneja la creación de usuarios y utiliza Long Polling
func CreateUsersHandler(c *gin.Context) {
	var usuario domain.User

	
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON"})
		return
	}

	
	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewCreateUser(repo)

	// Realizar la creación del usuario en un proceso que tomará tiempo
	// Ejemplo: proceso largo, por ejemplo, guardando en la base de datos
	go func() {
		// Simular una espera de 5 segundos, para representar un proceso largo
		time.Sleep(5 * time.Second)

		// Ejecutar el caso de uso para crear el usuario
		if err := useCase.Execute(usuario); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el usuario"})
			return
		}
	}()

	// Mantener la conexión abierta hasta que la tarea esté lista
	select {
	case <-time.After(10 * time.Second): 
		// Responder con un mensaje de éxito
		c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado con éxito"})
	case <-time.After(15 * time.Second):
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Tiempo de espera agotado"})
	}
}
