package main

import (
	"github.com/agilistikmal/jetship/app/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")

	// Register Routes
	routes.PackageRoutes(api)

	app.Listen(":9998")
}
