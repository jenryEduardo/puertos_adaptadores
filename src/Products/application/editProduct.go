package application

import "ejemplo/practica/src/Products/domain"

type EditProduct struct {
	repo domain.Iproduct
}

func NewUpdateProduct(repo domain.Iproduct) *EditProduct{
	return &EditProduct{repo: repo}
}

func (cp *EditProduct) Execute(id int,product *domain.Product)error{
	return cp.repo.Update(id,product)
}