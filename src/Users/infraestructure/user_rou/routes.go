package routes

import (
	"net/http"
	"ejemplo/practica/src/Users/infraestructure/controllers"
)

func SetupRoutesUsers() {
	http.HandleFunc("/users", controllers.CreateUsersHandler)
	http.HandleFunc("/view-users", controllers.GetUserHandler)
	http.HandleFunc("/delete-users/", controllers.DeleteUserHandeler)
	http.HandleFunc("/update-users/", controllers.EditUserHandler)
}