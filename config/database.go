// config/database.go
package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://pratama:cjzMmK1k7BGQ8zhq@yoginara.vvsjt.mongodb.net/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Test koneksi
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	log.Println("Connected to MongoDB!")
}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("stepupDB").Collection(collectionName)
}
