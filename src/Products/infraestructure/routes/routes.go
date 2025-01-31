package routes

import (
	"net/http"
	"ejemplo/practica/src/Products/infraestructure/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/products", controllers.CreateProductHandler)
	http.HandleFunc("/view-products", controllers.GetProductHandler)
	http.HandleFunc("/delete-products/", controllers.DeleteProductHandeler)
	http.HandleFunc("/update-products/", controllers.UpdateProductHandler)
}