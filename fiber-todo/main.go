package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{ID: 1, Name: "Send an email update to the team: 9am today", Completed: false},
	{ID: 2, Name: "Call the design agency to finalize mockups: 1:00pm", Completed: false},
	{ID: 3, Name: "Touch base with recruiters about new role: Tuesday", Completed: false},
	{ID: 4, Name: "Meet with the engineering team: Thursday", Completed: false},
}

func main() {
	app := fiber.New()

	// Logger
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi, there! 👋")
	})

	SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1

	todos := v1.Group("/todos")

	todos.Get("/", GetTodos)
	todos.Post("/", CreateTodo)
	todos.Get("/:id", GetTodo)
	todos.Delete("/:id", DeleteTodo)
	todos.Patch("/:id", UpdateTodo)
}

func GetTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}

	var body request

	err := c.BodyParser(&body)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
		return err
	}

	todo := &Todo{
		ID:        len(todos) + 1,
		Name:      body.Name,
		Completed: false,
	}

	todos = append(todos, todo)

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func GetTodo(c *fiber.Ctx) error {
	paramsID := c.Params("id")

	id, err := strconv.Atoi(paramsID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse id",
		})
	}

	for _, todo := range todos {
		if todo.ID == id {
			return c.Status(fiber.StatusOK).JSON(todo)
		}
	}

	c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Record not found",
	})
	return nil
}

func DeleteTodo(c *fiber.Ctx) error {
	paramsID := c.Params("id")

	id, err := strconv.Atoi(paramsID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse id",
		})
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)

			c.Status(fiber.StatusNoContent)
			return nil
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Record not found",
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	type request struct {
		Name      *string `json:"name"`
		Completed *bool   `json:"completed"`
	}

	paramsID := c.Params("id")

	id, err := strconv.Atoi(paramsID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse ID",
		})
	}

	var body request

	err = c.BodyParser(&body)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
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
		c.Status(fiber.StatusNotFound)
		return nil
	}

	if body.Name != nil {
		todo.Name = *body.Name
	}

	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}
