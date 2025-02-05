package controllers

import (
	"ejemplo/practica/src/Users/application"
	"ejemplo/practica/src/Users/domain"
	"ejemplo/practica/src/Users/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// EditUserHandler maneja la actualización de un usuario por ID con Long Polling
func EditUserHandler(c *gin.Context) {
	// Obtener el ID de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Decodificar el JSON de la solicitud
	var updatedUser domain.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON"})
		return
	}

	// Inicializar repositorio y caso de uso
	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewEditUser(repo)

	// Mantener la conexión abierta durante el proceso
	go func() {
		// Simulamos una demora de 5 segundos (por ejemplo, para hacer validaciones, actualizaciones en base de datos, etc.)
		time.Sleep(5 * time.Second)

		// Ejecutar la actualización
		if err := useCase.Execute(id, &updatedUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el usuario"})
			return
		}
		// Responder después de la actualización exitosa
		c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado con éxito"})
	}()

	// Mantener la conexión abierta durante el proceso
	select {
	case <-time.After(10 * time.Second): // Espera de 10 segundos antes de que se agote el tiempo de espera
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Tiempo de espera agotado"})
	}
}
