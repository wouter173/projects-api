package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/wouter173/projects/handlers"
)

func main() {
	godotenv.Load()

	app := fiber.New()

	app.Get("/:id/meta", handlers.IDMetaHandler)
	app.Get("/:id", handlers.IDHandler)
	app.Get("/", handlers.AllHandler)

	app.Listen(":" + os.Getenv("PORT"))
}
