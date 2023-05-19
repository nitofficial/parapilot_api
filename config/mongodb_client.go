package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDBClient() (*mongo.Client, error) {
	// Set up the MongoDB client options
	log.Print("This is a log message")
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.8.2")

	// Connect to the MongoDB server
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Print("This is a log message")
		return nil, err
	}

	return client, nil
}
