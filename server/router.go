package server

import (
	"github.com/christianmahardhika/mocktestgolang/service"
	"github.com/gofiber/fiber/v2"
)

func InitializeRouter(r *fiber.App, serviceController service.Controller) *fiber.App {
	r.Get("/todos", serviceController.FindTodos)
	r.Get("/todo", serviceController.FindTodoDetail)
	r.Post("/todo", serviceController.SaveTodo)
	r.Delete("/todo", serviceController.DeleteTodo)

	return r
}
