package controllers

import (
	"ejemplo/practica/src/Users/application"
	"ejemplo/practica/src/Users/infraestructure"
	"ejemplo/practica/src/Users/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Extraer el ID del producto desde la URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "ID del producto requerido en la URL", http.StatusBadRequest)
		return
	}

	// Convertir el ID a entero
	userID, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "ID del producto inválido", http.StatusBadRequest)
		return
	}

	var id int = userID

	
	// Decodificar el JSON del cuerpo de la solicitud
	var updatedUser domain.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewEditUser(repo)

	usuario := useCase.Execute(id,&updatedUser)

	if(usuario!=nil){
		fmt.Println("ocurrio un error al actulizar el usuario")
		return
	}

	fmt.Println("usuario actualizado")

}