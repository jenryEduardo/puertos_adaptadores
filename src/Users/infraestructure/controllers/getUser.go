package controllers

import (
	"ejemplo/practica/src/Users/application"
	"ejemplo/practica/src/Users/infraestructure"
	"encoding/json"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewGetUser(repo)
	users, err := useCase.Execute()

	if err != nil {
		http.Error(w, "No se pudo obtener los usuarios", http.StatusInternalServerError)
		return
	}

	// Convertir a JSON
	userJson, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJson) 
}
