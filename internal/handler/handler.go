package handler

import (
	"ma-backend-training/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes function
func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/api/v1/signin", SignInHandler)
	app.Post("/api/v1/signup", SignUpHandler)

	// Protected routes
	api := app.Group("/api/v1", middleware.JWTMiddleware)
	api.Get("/users", GetUsers)
	api.Delete("/user/:user_id", DeleteUser)
	api.Post("/user/:user_id", UpdateUser)
	api.Post("/user/:user_id/password", ResetPassword)
	api.Post("/files/upload", UploadFile)
	api.Get("/files/:file_id", GetFiles)
	api.Get("/files", GetFiles)
	api.Delete("/file/:file_id", DeleteFile)

	// External API integration routes
	api.Get("/integration/post/:post_id", GetAPIHandler)
	api.Post("/integration/post", PostAPIHandler)
}
