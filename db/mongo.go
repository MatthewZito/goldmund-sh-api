package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
)

func InitMongoSession() *mgo.Session {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := os.Getenv("MONGO_URI")
	s, err := mgo.Dial(connStr)

	if err != nil {
		panic(err)
	}
	return s
}
