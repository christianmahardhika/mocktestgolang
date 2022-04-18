package main

import (
	"github.com/christianmahardhika/mocktestgolang/server"
	"github.com/gofiber/fiber/v2"
)

var FiberApp *fiber.App

func init() {
	FiberApp = server.InitiateServer()
}

func main() {
	server.StartApplication(FiberApp, "8080")
}
