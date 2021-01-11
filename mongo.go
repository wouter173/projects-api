package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//GetMongoConnection get mongodb connection from env DBURI
func GetMongoConnection() (*mongo.Client, error) {

	options := options.Client().ApplyURI(os.Getenv("DBURI"))

	client, err := mongo.Connect(context.Background(), options)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		return nil, err
	}

	return client, nil
}

//GetMongoCollection from dbName and collectionName
func GetMongoCollection(dbName string, collectionName string) (*mongo.Collection, error) {

	client, err := GetMongoConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)

	return collection, nil
}
