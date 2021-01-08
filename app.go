package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("DBURI"))

	connect()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DBURI")))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	defer client.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	databases, _ := client.ListDatabaseNames(ctx, bson.M{})
	fmt.Println(databases)

	database := client.Database("Website")
	collections, _ := database.ListCollectionNames(ctx, bson.M{})
	fmt.Println(collections)

	collection := database.Collection("Projects")

	filter := bson.D{{"name", "Shrt.lu"}}

	var result Project

	err = collection.FindOne(ctx, filter).Decode(&result)

	fmt.Printf("%+v\n", result)
	fmt.Println(collection.EstimatedDocumentCount(ctx))
}
