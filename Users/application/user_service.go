package application

import (
	"ejemplo/practica/Users/domain"
	"fmt"
)

//llamamos al caso de uso para poder crear

type UserService struct {
	userRepo domain.UserRepository
}

func NewRepositoryService(repo domain.UserRepository) *UserService{
	return &UserService{userRepo: repo}
}

func (s *UserService) CrearUser(user domain.User)error {
	if user.Nombre == "" {
		return fmt.Errorf("necesita ingresar un nombre")
	}
	return s.CrearUser(user)
}