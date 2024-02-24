package routes

import (
	"github.com/agilistikmal/jetship/app/controller"
	"github.com/gofiber/fiber/v2"
)

func PackageRoutes(api fiber.Router) {
	route := api.Group("/package")

	route.Get("/", controller.FindPackage)
	route.Get("/:id", controller.FindPackage)
}
