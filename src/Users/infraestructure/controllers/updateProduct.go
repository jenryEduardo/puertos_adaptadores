package controllers

import (
	"encoding/json"
	"ejemplo/practica/src/Users/application"
	"ejemplo/practica/src/Users/domain"
	"ejemplo/practica/src/Users/infraestructure"
	"net/http"
	"strconv"
	"strings"
)

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	// Verificar que el método es PUT
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Extraer el ID del producto desde la URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "ID del producto requerido en la URL", http.StatusBadRequest)
		return
	}

	// Convertir el ID a entero
	productID, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "ID del producto inválido", http.StatusBadRequest)
		return
	}

	print(productID)

	// Decodificar el JSON del cuerpo de la solicitud
	var updatedProduct domain.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	// Crear repositorio e instancia del caso de uso
	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewUpdateProduct(repo)

	// Ejecutar la actualización
	err = useCase.Execute(productID, &updatedProduct)
	if err != nil {
		http.Error(w, "Error al actualizar el producto", http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Producto actualizado correctamente"))
}
