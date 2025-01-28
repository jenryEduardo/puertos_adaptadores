package helpers

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type MongoDBHelper struct {
	Client *mongo.Client
}


func NewMongoDBHelper(uri string) *MongoDBHelper {
	clientOpcions := options.Client().ApplyURI(uri)


	Client, err:= mongo.Connect(context.TODO(),clientOpcions)

	if err != nil {
		log.Fatal("error al conectarse a mongo",err)
	}

	if err := Client.Ping(context.TODO(),nil);err !=nil {
		log.Fatal("no se pudo conectar a la base de datos")
	}

	return &MongoDBHelper{Client: Client}
}

func (h *MongoDBHelper) GetCollection(databaseName, collectionName string) *mongo.Collection {
	return h.Client.Database(databaseName).Collection(collectionName)
}



// CloseConnection cierra la conexión con MongoDB
func (h *MongoDBHelper) CloseConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := h.Client.Disconnect(ctx); err != nil {
		log.Fatalf("Error al cerrar la conexión con MongoDB: %v", err)
	}

	log.Println("Conexión a MongoDB cerrada correctamente")
}