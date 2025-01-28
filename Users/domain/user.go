package domain
//definimos el modelo para el usuario donde sera parte del caso de uso
type User struct {
	Nombre string
	Correo string
	Password string
	Num_tel int16
}

//crear un archivo donde se almacenen los repository

//definimos el repositorio del modelo (puertos)
