package domain

type IUser interface {
	Save(user *User)error
	GetAll()([]User,error)
	Update(id int,user *User)error
	Delete(id int)error
}