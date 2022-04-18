package server

import (
	"github.com/christianmahardhika/mocktestgolang/service"
	"github.com/gofiber/fiber/v2"
)

func InitiateServer() *fiber.App {
	// Initialize the router
	r := fiber.New()

	// Initialize the usecase
	serviceUC := service.NewUseCase(service.NewRepository())

	// Initialize the controller
	serviceController := service.Controller{UseCase: serviceUC}

	// Initialize the routes
	r = InitializeRouter(r, serviceController)
	return r

}
