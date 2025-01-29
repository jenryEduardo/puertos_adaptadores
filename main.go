package main

import (
	"log"
	"net/http"
	"ejemplo/practica/src/Users/infraestructure/routes"
)

func main() {
	routes.SetupRoutes()
	log.Println("Servidor escuchando en puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}