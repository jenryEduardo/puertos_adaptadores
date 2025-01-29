package domain


type Iproduct interface {
	Save(product *Product)error
	GetAll()([]Product,error)
}