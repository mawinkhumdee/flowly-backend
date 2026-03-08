package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mawinkhumdee/flowly-project/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use URI from config.yml
	clientOptions := options.Client().ApplyURI(config.AppConfig.Database.URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	Client = client
}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database(config.AppConfig.Database.Name).Collection(collectionName)
}
