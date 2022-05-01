package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func StartApplication(r *fiber.App, port string) {

	// Start the server
	err := r.Listen(":" + port)
	if err != nil {
		panic(err)
	}

}

func ShutdownApplication(r *fiber.App) {

	// Implement graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		_ = <-ctx.Done()
		log.Println("Gracefully shutting down...")
		_ = r.Shutdown()
		stop()
	}()
}

func CleanUpApplication() {
	log.Println("Running cleanup tasks...")
	// wait 2 seconds for the server to shutdown
	time.Sleep(2 * time.Second)
	context.WithTimeout(context.Background(), 5*time.Second)
	DBCloseConnection()
	log.Println("Finish cleanup tasks...")
}
