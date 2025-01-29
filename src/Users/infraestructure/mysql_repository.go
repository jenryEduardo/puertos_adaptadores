package infraestructure

import (
	"ejemplo/practica/src/core"
	"ejemplo/practica/src/Users/domain"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}

func NewMySQLRepository() *MySQLRepository {
	conn := core.GetDBPool()
	return &MySQLRepository{conn: conn}
}

func (r *MySQLRepository) Save(p *domain.Product) error {
	query := "INSERT INTO Product (nombre, precio) VALUES (?, ?)"
	_, err := r.conn.DB.Exec(query, p.Nombre, p.Precio)
	return err
}

func (r *MySQLRepository) GetAll() ([]domain.Product, error) {
	query := "SELECT nombre, precio FROM Product"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Nombre, &product.Precio); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}