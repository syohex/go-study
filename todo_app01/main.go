package main

import (
	"todoapp001/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the end point",
		})
	})

	api := app.Group("/api")
	api.Get("", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Your are at the api endpoint",
		})
	})

	route.TodoRoute(api.Group("/todos"))
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
