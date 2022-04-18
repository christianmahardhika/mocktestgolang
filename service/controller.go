package service

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	UseCase UseCase
}

func (c *Controller) SaveTodo(ctx *fiber.Ctx) error {
	todoAll := &TodoAll{}
	if err := ctx.BodyParser(todoAll); err != nil {
		return ctx.Status(400).JSON("bad request")
	}
	res, err := c.UseCase.SaveTodo(todoAll)
	if err != nil {
		return ctx.Status(500).JSON("internal server error")
	}
	return ctx.Status(201).JSON(res)
}

func (c *Controller) FindTodos(ctx *fiber.Ctx) error {
	res, err := c.UseCase.GetTodos()
	if err != nil {
		return ctx.Status(500).JSON("internal server error")
	}
	return ctx.Status(201).JSON(res)
}

func (c *Controller) FindTodoDetail(ctx *fiber.Ctx) error {

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON("bad request")
	}
	res, err := c.UseCase.GetTodoDetail(id)
	if err != nil {
		return ctx.Status(500).JSON("internal server error")
	}
	return ctx.Status(201).JSON(res)
}

func (c *Controller) DeleteTodo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON("bad request")
	}
	err = c.UseCase.DeleteTodo(id)
	if err != nil {
		return ctx.Status(500).JSON("internal server error")
	}
	return ctx.Status(201).JSON("success")
}