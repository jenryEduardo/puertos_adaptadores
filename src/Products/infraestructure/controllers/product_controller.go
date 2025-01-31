package controllers


import (
	"encoding/json"
	"net/http"
	"ejemplo/practica/src/Products/application"
	"ejemplo/practica/src/Products/domain"
	"ejemplo/practica/src/Products/infraestructure"
)


func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var product domain.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
		return
	}
	//inicializamos la bd
	repo := infraestructure.NewMySQLRepository()
	//se crea el caso de uso
	useCase := application.NewCreateProduct(repo)
	//se llama al metodo execute y si falla mandamos un mensaje
	if err := useCase.Execute(product); err != nil {
		http.Error(w, "Error al guardar el producto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Producto creado con éxito"))
}
