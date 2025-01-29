package controllers

import (
	"ejemplo/practica/src/Users/application"
	"ejemplo/practica/src/Users/infraestructure"
	// "encoding/json"
	"net/http"
	"strings"
)


func DeleteProductHandeler(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")

	// Validar que haya un nombre después de "/delete-products/"
	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "Nombre del producto requerido", http.StatusBadRequest)
		return
	}

	productName := pathParts[2] // Obtener el nombre del producto


	var Nombreproduct string = productName
	// if err := json.NewDecoder(r.Body).Decode(&Nombreproduct); err != nil {
	// 	http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
	// 	return
	// }

	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewDeleteProduct(repo)

	if err := useCase.Execute(Nombreproduct); err != nil {
		http.Error(w, "Error al guardar el producto", http.StatusInternalServerError)
		return
}
}