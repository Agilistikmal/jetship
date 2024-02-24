package controller

import (
	"github.com/gofiber/fiber/v2"
)

func FindPackages(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"hello": "world",
	})
}

func FindPackage(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"hello": "world",
	})
}
