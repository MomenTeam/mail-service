package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Context context.Context
var CancelFunc context.CancelFunc
var MailTemplatesCollection *mongo.Collection

func Setup() {
	Client, Context, CancelFunc = getConnection(os.Getenv("DB_CONNECTION"))

	getCollections()
}

func getConnection(connectionString string) (*mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	return client, ctx, cancel
}

func getCollections() {
	databaseName := "hackathon"
	database := Client.Database(databaseName)

	MailTemplatesCollection = database.Collection("MailTemplates")
}
