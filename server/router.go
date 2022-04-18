package server

import (
	"github.com/christianmahardhika/mocktestgolang/service"
	"github.com/gofiber/fiber/v2"
)

func InitializeRouter(r *fiber.App, serviceController service.Controller) *fiber.App {
	r.Get("/todos", serviceController.FindTodos)
	r.Get("/todos/:id", serviceController.FindTodoDetail)
	r.Post("/todos", serviceController.SaveTodo)
	r.Delete("/todos/:id", serviceController.DeleteTodo)

	return r
}
