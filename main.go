package main

import (
	"log"
	"net/http"
	productRoutes "ejemplo/practica/src/Products/infraestructure/routes" // Renombrado
	userRoutes "ejemplo/practica/src/Users/infraestructure/user_rou"     // Renombrado
)

func main(){
	productRoutes.SetupRoutes()
	userRoutes.SetupRoutesUsers()
	log.Println("Servidor escuchando en puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
