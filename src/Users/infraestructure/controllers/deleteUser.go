package controllers

import (
	"ejemplo/practica/src/Users/application"
	"ejemplo/practica/src/Users/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// DeleteUserHandler maneja la eliminación de un usuario con Long Polling
func DeleteUserHandler(c *gin.Context) {
	// Obtener el ID de la URL
	idUser := c.Param("id")

	// Convertir el ID a entero
	id, err := strconv.Atoi(idUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	// Inicializar repositorio y caso de uso
	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewDeleteProduct(repo)

	// Mantener la conexión abierta durante el proceso
	go func() {
		// Simular un proceso largo de eliminación, por ejemplo, 5 segundos de espera
		time.Sleep(5 * time.Second)

		// Ejecutar la lógica para eliminar el usuario
		if err := useCase.Execute(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario"})
			return
		}
		// Responder después de la eliminación exitosa
		c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado con éxito"})
	}()

	// Mantener la conexión abierta durante la operación
	select {
	case <-time.After(10 * time.Second): // Esperar hasta 10 segundos
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Tiempo de espera agotado"})
	}
}
