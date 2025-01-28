package infraestructure

import (
	"context"
	"ejemplo/practica/Users/domain"
	"ejemplo/practica/helpers"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	Collection *mongo.Collection
}

// NewMongoUserRepository inicializa el repositorio con MongoDB
func NewMongoUserRepository() *MongoUserRepository {
	// Crear instancia del helper y obtener la colección
	mongoHelper := helpers.NewMongoDBHelper("mongodb://localhost:27017")
	collection := mongoHelper.GetCollection("tienda", "users")

	return &MongoUserRepository{Collection: collection}
}

func (r *MongoUserRepository) Save(u *domain.User)error{
	_,erro := r.Collection.InsertOne(context.TODO(),u)
	return erro
}