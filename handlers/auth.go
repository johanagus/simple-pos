package handlers

import "github.com/gofiber/fiber/v2"

func Auth(c *fiber.Ctx) error {
	return c.SendString("Hallo ini adalah auth")
}
