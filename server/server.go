package server

import (
	"github.com/christianmahardhika/mocktestgolang/service"
	"github.com/gofiber/fiber/v2"
)

func InitiateServer(dbString string, dbName string) *fiber.App {
	// Initialize the router
	r := fiber.New()

	// Initialize the database
	dbConnnection := GetDBConnection(dbString, dbName)
	// Initialize the usecase
	serviceUC := service.NewUseCase(service.NewRepository(dbConnnection))

	// Initialize the controller
	serviceController := service.Controller{UseCase: serviceUC}

	// Initialize the routes
	r = InitializeRouter(r, serviceController)
	return r

}
