package main

import (
	"github.com/christianmahardhika/mocktestgolang/server"
	"github.com/gofiber/fiber/v2"
)

var FiberApp *fiber.App

func init() {

	dbString := "mongodb://root:root@localhost:27017"
	dbName := "mocktestgolang"
	FiberApp = server.InitiateServer(dbString, dbName)
}

func main() {
	port := "8080"

	server.ShutdownApplication(FiberApp)

	server.StartApplication(FiberApp, port)

	server.CleanUpApplication()
}
