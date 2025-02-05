package controllers

import (
	"ejemplo/practica/src/Products/application"
	"ejemplo/practica/src/Products/domain"
	"ejemplo/practica/src/Products/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"

)

// CreateProductHandlerChunked - Long Polling para creación de productos
func CreateProductHandler(c *gin.Context) {
	var product domain.Product

	// Decodificar el JSON del cuerpo de la solicitud
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON"})
		return
	}

	// Inicializar la BD y el caso de uso
	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewCreateProduct(repo)

	// Ejecutar la lógica de creación del producto
	if err := useCase.Execute(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el producto"})
		return
	}

	// Devolver respuesta JSON estándar
	c.JSON(http.StatusOK, gin.H{"message": "Producto creado con éxito"})
}
