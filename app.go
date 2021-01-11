package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New()

	app.Get("/:id", IDHandler)
	app.Get("/", AllHandler)

	app.Listen(":" + os.Getenv("PORT"))
}
