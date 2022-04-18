package server

import "github.com/gofiber/fiber/v2"

func StartApplication(r *fiber.App, port string) {

	// Start the server
	err := r.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
