package main

import (
	"github.com/Salavei/golang_restapi/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
