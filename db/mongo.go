package db

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMongoSession initializes a session to the environmentally-specified cluster and returns a pointer to the similarly specified collection
func InitMongoSession() (*mongo.Collection, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	endpoint := os.Getenv("MONGO_URI")

	clientOptions := options.Client().ApplyURI(endpoint)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database("test").Collection("entries")

	return collection, nil
}
