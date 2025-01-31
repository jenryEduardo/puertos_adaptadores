package application

import "ejemplo/practica/src/Products/domain"
//se crea el servicio hacia el domain
type GetProduct struct {
	repo domain.Iproduct
}	

//nos sirve para cuando sea invocado en infraestructure (controllers) creee el caso de uso 
//en pocas palabras como decirle a infra lo que puede hacer
func NewGetProduct(repo domain.Iproduct) *GetProduct{
		return &GetProduct{repo: repo}
}
//metodo que sirve para llamar a un metodo que realice una operacion desde x lugar 
//ya que no depende de nadie
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

