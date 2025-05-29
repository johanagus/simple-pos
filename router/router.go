package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-pos/handlers"
)

func New(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api/v1")

	api.Post("/auth", handlers.Auth)

	// api.Get("/users")
	// api.Get("/user/:id")
	// api.Post("/user")
	// api.Put("/user/:id")
	// api.Delete("/user/id:

	// api.Get("/products")
	// api.Get("/product/:id")
	// api.Post("/product")
	// api.Put("/product/:id")
	// api.Delete("/product/:id

	// api.Get("/categories")
	// api.Get("/category/:id")
	// api.Post("/category")
	// api.Put("/category/:id")
	// api.Delete("/category/:id")

}
