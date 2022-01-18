package mongo_db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongo_db.org/mongo_db-driver/mongo_db/readpref"
)

const (
	mongodbUri = "MONGODB_URI"
)

var (
	DB  *mongo.Database
	uri string
)

func loadEnvironment() {
	err := godotenv.Load("./datasources/.env")

	if err != nil {
		log.Println("Couldn't load environment variables")
		panic(err.Error())
	}
}

func init() {
	var err error
	loadEnvironment()
	uri = os.Getenv(mongodbUri)

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("test")

	fmt.Println("Connected to MongoDB!")
}
