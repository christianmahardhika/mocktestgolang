package server

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbConn *mongo.Database

func InitDB(dbString string, dbName string) {
	// Initialize the database
	var ctx = context.Background()
	clientOptions := options.Client()
	clientOptions.ApplyURI(dbString)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	dbConn = client.Database(dbName)
}

func GetDBConnection(dbString string, dbName string) *mongo.Database {
	if dbConn == nil {
		InitDB(dbString, dbName)
		return dbConn
	} else {
		return dbConn
	}
}
