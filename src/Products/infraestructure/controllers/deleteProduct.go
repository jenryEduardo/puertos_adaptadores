package controllers

import (
	"ejemplo/practica/src/Products/application"
	"ejemplo/practica/src/Products/infraestructure"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// DeleteProductHandlerChunked - Long Polling para eliminación de productos
func DeleteProductHandler(c *gin.Context) {
	productName := c.Param("name")

	if productName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre del producto requerido"})
		return
	}

	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewDeleteProduct(repo)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	encoder := json.NewEncoder(c.Writer)

	// Notificación inicial
	encoder.Encode(gin.H{"status": "Recibiendo solicitud para eliminar producto"})
	c.Writer.Flush()
	time.Sleep(2 * time.Second)

	// Intento de eliminación
	if err := useCase.Execute(productName); err != nil {
		encoder.Encode(gin.H{"error": "Error al eliminar el producto"})
		c.Writer.Flush()
		return
	}

	// Confirmación de eliminación
	encoder.Encode(gin.H{"message": "Producto eliminado correctamente"})
	c.Writer.Flush()
}
