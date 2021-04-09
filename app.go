package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/wouter173/projects/handlers"
)

func main() {
	godotenv.Load()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://wouterdb.nl, https://wouter173.nl",
	}))

	app.Get("/:id/meta", handlers.IDMetaHandler)
	app.Get("/:id", handlers.IDHandler)
	app.Get("/", handlers.AllHandler)

	app.Listen(":" + os.Getenv("PORT"))
}
