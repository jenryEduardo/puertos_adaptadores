package controllers

import (
	"ejemplo/practica/src/Users/application"
	"ejemplo/practica/src/Users/infraestructure"
	"net/http"
	"strconv"
	"strings"
	"fmt"
)


func DeleteUserHandeler(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")

	// Validar que haya un nombre después de "/delete-products/"
	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "id del usuario requerido", http.StatusBadRequest)
		return
	}

	idUser := pathParts[2] // Obtener el nombre del producto
	id,err :=  strconv.Atoi(idUser)

	if err != nil {
		fmt.Println("Error: No se pudo convertir a entero:", err)
		return
	}

	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewDeleteProduct(repo)

	if err := useCase.Execute(id); err != nil {
		http.Error(w, "Error al eliminar el producto", http.StatusInternalServerError)
		return
}
}