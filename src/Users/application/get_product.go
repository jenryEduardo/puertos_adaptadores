package application

import "ejemplo/practica/src/Users/domain"

type GetProduct struct {
	repo domain.Iproduct
}


func NewGetProduct(repo domain.Iproduct) *GetProduct{
		return &GetProduct{repo: repo}
}

func (cp *GetProduct) Execute() ([]domain.Product, error) {
	return cp.repo.GetAll()
}


// func (cp *GetProduct) Execute() ([]domain.Product, error) {
// 	// Llama al método GetAll del repositorio
// 	products, err := cp.repo.GetAll()
// 	if err != nil {
// 		return nil, err // Devuelve el error si ocurre
// 	}

// 	return products, nil // Devuelve los productos y ningún error
// }