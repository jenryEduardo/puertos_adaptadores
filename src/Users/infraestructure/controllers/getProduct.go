package controllers

import (
	"ejemplo/practica/src/Users/application"
	"ejemplo/practica/src/Users/infraestructure"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	
	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewGetProduct(repo)
	products, err := useCase.Execute()
	if err != nil {
		fmt.Printf("error al obtner los productos", err)
		return
	}

	productsJson, err := json.Marshal(products)

	w.WriteHeader(http.StatusCreated)
	fmt.Printf(string(productsJson))
}
