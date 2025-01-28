package application

import (
	"ejemplo/practica/Users/domain"
)

//llamamos al repositorio del dominio para poder hacer uso de ellas

type UserService struct {
	//creamos una variable de tipo user repository
	userRepo domain.UserRepository
}

// creo una variable de tipo repo que sera igual al repositorio con el puntero señalando 
//hacia el dominio para despues retornar el repositorio ya inicializado

func NewRepositoryService(repo domain.UserRepository) *UserService{
	return &UserService{userRepo: repo}
}
//ya que tenemos inicializado el struc de user creamos una var del tipo userService y el metodo de crear
//basado en el modelo 
func (s *UserService) CrearUser(user domain.User) {

	 s.userRepo.Save(&user)
}

//ver y preguntar que debe devolver el return