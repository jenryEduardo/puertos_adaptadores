package domain
//definimos el modelo para el usuario donde sera parte del caso de uso
type User struct {
	Nombre string
	Correo string
	Password string
	Num_tel int16
}

//definimos el repositorio del modelo (puertos)
type UserRepository interface {
	save(user *User) error
	findById(id string)(*User,error)
	findAll()([]*User,error)
}