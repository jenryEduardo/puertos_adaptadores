package routes

import (
	"net/http"
	"ejemplo/practica/src/Users/infraestructure/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/products", controllers.CreateProductHandler)
	http.HandleFunc("/productos", controllers.GetProductHandler)
}