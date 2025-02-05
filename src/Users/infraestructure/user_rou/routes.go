package routes

import (
	"ejemplo/practica/src/Users/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutesUsers configura las rutas para usuarios
func SetupRoutesUsers(router *gin.Engine) {

	routes:=router.Group("/users")

	{
		routes.POST("/", controllers.CreateUsersHandler)   // Crear usuario
		routes.GET("/", controllers.GetUserHandler)        // Obtener usuarios
		routes.DELETE("/:id", controllers.DeleteUserHandler) // Eliminar usuario por ID
		routes.PUT("/:id", controllers.EditUserHandler)    // Actualizar usuario por ID
	}
}
