package controllers

import (
	"ejemplo/practica/src/Products/application"
	"ejemplo/practica/src/Products/domain"
	"ejemplo/practica/src/Products/infraestructure"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// UpdateProductHandlerChunked - Long Polling para actualización de productos
func UpdateProductHandler(c *gin.Context) {
	// Obtener el ID del producto desde los parámetros de la URL
	productIDStr := c.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto inválido"})
		return
	}

	var updatedProduct domain.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar el JSON"})
		return
	}

	
	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewUpdateProduct(repo)

	// Configurar encabezados para chunked
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	encoder := json.NewEncoder(c.Writer)

	// Enviar primer chunk indicando que la actualización ha comenzado
	encoder.Encode(gin.H{"status": "Iniciando la actualización del producto"})
	c.Writer.Flush()
	time.Sleep(1 * time.Second) 

	
	if err := useCase.Execute(productID, &updatedProduct); err != nil {
		encoder.Encode(gin.H{"error": "Error al actualizar el producto"})
		c.Writer.Flush()
		return
	}


	encoder.Encode(gin.H{"message": "Producto actualizado correctamente"})
	c.Writer.Flush()
}
