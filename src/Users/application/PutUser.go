package application

import "ejemplo/practica/src/Users/domain"

type EditUser struct {
	repo domain.IUser
}

func NewEditUser(repo domain.IUser)*EditUser{
	return &EditUser{repo: repo}
}

func (cp *EditUser) Execute(id int,p *domain.User)error{
	return cp.repo.Update(id,p)
}
