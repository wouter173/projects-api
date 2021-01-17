package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

//AllHandler get all projects from db
func AllHandler(c *fiber.Ctx) error {
	c.Set("content-type", "application/json")
	c.Set("Access-Control-Allow-Origin", "*")

	collection, err := GetMongoCollection(os.Getenv("DB"), os.Getenv("COLL"))

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	filter := bson.M{}
	var results []Project

	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	cur.All(context.Background(), &results)

	if results == nil {
		return c.Status(404).JSON("Not found")
	}

	json, _ := json.Marshal(results)

	return c.Send(json)
}

//IDHandler get 1 project from db based on id
func IDHandler(c *fiber.Ctx) error {
	c.Set("content-type", "application/json")
	c.Set("Access-Control-Allow-Origin", "*")

	collection, err := GetMongoCollection(os.Getenv("DB"), os.Getenv("COLL"))

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	var filter bson.M

	if c.Params("id") != "" {
		filter = bson.M{"id": c.Params("id")}
	}

	var result Project

	collection.FindOne(context.Background(), filter).Decode(&result)

	if result.ID == "" {
		return c.Status(404).JSON("Not found")
	}

	json, _ := json.Marshal(result)

	return c.Send(json)
}
