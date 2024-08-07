package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"ma-backend-training/config"
	_ "ma-backend-training/docs" // Import Swagger docs
	"ma-backend-training/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title Fiber Swagger Example API
// @version 1.0
// @description This is a sample server for demonstrating Swagger with Fiber
// @host localhost:3000
// @BasePath /api/v1

func main() {
	app := fiber.New()

	// Load configuration
	config.LoadConfig()

	// Connect to MongoDB
	config.ConnectMongoDB()

	// Middleware
	app.Use(logger.New())

	// Swagger route
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// Register routes
	handler.RegisterRoutes(app)

	// Create a channel to listen for signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Run the server in a goroutine
	go func() {
		if err := app.Listen(":" + config.AppConfig.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Block until we receive a signal
	<-quit

	log.Println("Shutting down server...")

	// Shut down the server gracefully
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
