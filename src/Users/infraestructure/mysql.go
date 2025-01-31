package infraestructure

import (
	"ejemplo/practica/src/Users/domain"
	"ejemplo/practica/src/core"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}



func NewMySQLRepository() *MySQLRepository {
	conn := core.GetDBPool()
	return &MySQLRepository{conn: conn}
}

func (r *MySQLRepository) Save(p *domain.User) error {
	query := "INSERT INTO Users (nombre, correo,contraseña) VALUES (?, ?,?)"
	_, err := r.conn.DB.Exec(query, p.Nombre, p.Correo,p.Contraseña)
	return err
}

func (r *MySQLRepository) Delete(p int)error{
	id:=p
	query := "DELETE FROM Users WHERE id = ?"
	_,err :=r.conn.DB.Exec(query,id)
	return err
}

func (r *MySQLRepository) Update(id int,p *domain.User)error{
	query := "UPDATE Users SET nombre = ?, correo = ?, contraseña=? WHERE id = ?"
    _, err := r.conn.DB.Exec(query, p.Nombre, p.Correo,p.Contraseña,id)
    if err != nil {
        return err
    }
	return err
}

func (r *MySQLRepository) GetAll() ([]domain.User, error) {
	query := "SELECT nombre, correo FROM Users"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.Nombre, &user.Correo); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
