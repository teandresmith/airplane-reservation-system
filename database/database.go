package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func CreateMongoClient() *mongo.Client{
	MongoURI := os.Getenv("MONGO_URI")
	if MongoURI == "" {
		err := godotenv.Load()
		if err != nil {
			log.Panic(err)
		}
		MongoURI = os.Getenv("MONGO_URI")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI))
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	connectErr := client.Connect(ctx)
	if connectErr != nil {
		log.Panic(connectErr)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

var Client *mongo.Client = CreateMongoClient() 

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("SkyHigh").Collection(collectionName)
	return collection
}


