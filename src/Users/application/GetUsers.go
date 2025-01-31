package application

import "ejemplo/practica/src/Users/domain"

type GetUser struct {
	repo domain.IUser
}

func NewGetUser(repo domain.IUser) *GetUser{
	return &GetUser{repo: repo}
}

func (cp *GetUser) Execute()([]domain.User,error)  {
	return cp.repo.GetAll()
}