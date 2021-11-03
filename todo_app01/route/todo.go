package route

import (
	"todoapp001/controller"

	"github.com/gofiber/fiber/v2"
)

func TodoRoute(router fiber.Router) {
	router.Get("", controller.GetTodos)
	router.Post("", controller.CreateTodo)
	router.Put("/:id", controller.UpdateTodo)
	router.Delete("/:id", controller.DeleteTodo)
	router.Get("/:id", controller.GetTodo)
}
