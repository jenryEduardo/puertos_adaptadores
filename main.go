package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	productRoutes "ejemplo/practica/src/Products/infraestructure/routes"
	userRoutes "ejemplo/practica/src/Users/infraestructure/user_rou"
)

func main() {
	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Permitir Angular en puerto 4200
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Cache de preflight request
	}))

	// Configurar rutas de productos y usuarios
	productRoutes.SetupRoutes(router)
	userRoutes.SetupRoutesUsers(router)

	port := ":8080"
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(router.Run(port))
}
