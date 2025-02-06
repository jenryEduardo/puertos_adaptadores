package application

import (
	"ejemplo/practica/src/Products/domain"
)

//realizamos una inyeccion de dependencias que nos permita usar el contrato(metodos) creado en domain
type CreateProduct struct {
	repo domain.Iproduct
}


func NewCreateProduct(repo domain.Iproduct) *CreateProduct {
	return &CreateProduct{repo: repo}
}

func (cp *CreateProduct) Execute(p domain.Product)error{
	return cp.repo.Save(&p)
}


//realizar una infografia de esos 3 modelos
//mvc mvp mvm