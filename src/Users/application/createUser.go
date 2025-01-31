package application

import "ejemplo/practica/src/Users/domain"

type CreateUser struct {
	repo domain.IUser
}


func NewCreateUser(repo domain.IUser)*CreateUser {
	return &CreateUser{repo: repo}
}


func (cp *CreateUser) Execute(p domain.User)error{
	return cp.repo.Save(&p)
}