package controllers

import (
	"ejemplo/practica/src/Products/application"
	"ejemplo/practica/src/Products/infraestructure"

	"github.com/gin-gonic/gin"
	"net/http"

)

func GetProductHandler(c *gin.Context) {
	// Configurar la cabecera correctamente
	c.Writer.Header().Set("Content-Type", "application/json")

	// Obtener productos desde la base de datos
	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewGetProduct(repo)

	products, err := useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener los productos"})
		return
	}

	// Enviar la respuesta como un JSON completo
	c.JSON(http.StatusOK, products)
}
