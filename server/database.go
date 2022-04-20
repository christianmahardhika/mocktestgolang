package server

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbConn *mongo.Database

func InitDB() {
	// Initialize the database
	var ctx = context.Background()
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://root:root@localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	dbConn = client.Database("mocktestgolang")
}

func GetDBConnection() *mongo.Database {
	if dbConn == nil {
		InitDB()
		return dbConn
	} else {
		return dbConn
	}
}
