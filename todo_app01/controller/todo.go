package controller

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{
		ID:        1,
		Title:     "Walk the dog",
		Completed: false,
	},
	{
		ID:        2,
		Title:     "Walk the cat",
		Completed: false,
	},
}

func GetTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": todos,
		},
	})
}

func CreateTodo(c *fiber.Ctx) error {
	type Request struct {
		Title string `json:"title"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse request JSON",
		})
	}

	todo := &Todo{
		ID:        int64(len(todos) + 1),
		Title:     body.Title,
		Completed: false,
	}
	todos = append(todos, todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

func GetTodo(ctx *fiber.Ctx) error {
	paramID := ctx.Params("id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse ID",
		})
	}

	for _, todo := range todos {
		if todo.ID == id {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": fiber.Map{
					"todo": todo,
				},
			})
		}
	}

	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Todo not found",
	})
}

func UpdateTodo(ctx *fiber.Ctx) error {
	paramID := ctx.Params("id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse ID",
		})
	}

	type Request struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
	}

	var body Request
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	var todo *Todo
	for _, t := range todos {
		if t.ID == id {
			todo = t
			break
		}
	}

	if todo == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not found",
		})
	}

	if body.Title != nil {
		todo.Title = *body.Title
	}
	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

func DeleteTodo(ctx *fiber.Ctx) error {
	paramID := ctx.Params("id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parser ID",
		})
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
				"success": true,
				"message": "Deleted Successfully",
			})
		}
	}

	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Todo not found",
	})
}
