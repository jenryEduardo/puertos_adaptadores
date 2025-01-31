package application

import "ejemplo/practica/src/Users/domain"

type DeleteUser struct {
	repo domain.IUser
}

func NewDeleteProduct(repo domain.IUser) *DeleteUser {
	return &DeleteUser{repo: repo}
}

func (cp *DeleteUser) Execute(id int)error{
	return cp.repo.Delete(id)
}