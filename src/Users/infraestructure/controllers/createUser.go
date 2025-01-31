package controllers


import (
	"encoding/json"
	"net/http"
	"ejemplo/practica/src/Users/application"
	"ejemplo/practica/src/Users/domain"
	"ejemplo/practica/src/Users/infraestructure"
)


func CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var usuario domain.User
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
		return
	}
	//inicializamos la bd
	repo := infraestructure.NewMySQLRepository()
	//se crea el caso de uso
	useCase := application.NewCreateUser(repo)
	//se llama al metodo execute y si falla mandamos un mensaje
	if err := useCase.Execute(usuario); err != nil {
		http.Error(w, "Error al guardar el usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("usuario creado con éxito"))
}
