package routes

import (
	"ejemplo/practica/src/Products/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configura las rutas con Gin
func SetupRoutes(router *gin.Engine) {


	product := router.Group("/products")
	// Rutas para productos
	product.POST("/", controllers.CreateProductHandler)
	product.GET("/", controllers.GetProductHandler)
	product.DELETE("/:name", controllers.DeleteProductHandler)
	product.PUT("/:id", controllers.UpdateProductHandler)
}
