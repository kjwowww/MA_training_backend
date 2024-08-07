package handler

import (
	"ma-backend-training/internal/service"

	"github.com/gofiber/fiber/v2"
)

// SignUpHandler handles user registration
func SignUpHandler(c *fiber.Ctx) error {
	type request struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"status":  false,
			"message": "Invalid request payload",
			"data":    nil,
		})
	}

	if req.FirstName == "" || req.LastName == "" || req.Username == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"status":  false,
			"message": "Missing required fields",
			"data":    nil,
		})
	}

	err := service.CreateUser(req.FirstName, req.LastName, req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"status":  false,
			"message": "Failed to create user",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    201,
		"status":  true,
		"message": "User created successfully",
		"data":    nil,
	})
}

// SignInHandler handles user authentication
func SignInHandler(c *fiber.Ctx) error {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"status":  false,
			"message": "Invalid request payload",
			"data":    nil,
		})
	}

	user, err := service.GetUserByUsername(req.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    401,
			"status":  false,
			"message": "Invalid username or password",
			"data":    nil,
		})
	}

	if !service.CheckPassword(req.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    401,
			"status":  false,
			"message": "Invalid username or password",
			"data":    nil,
		})
	}

	token, err := service.GenerateJWT(user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"status":  false,
			"message": "Failed to generate token",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"status":  true,
		"message": "Sign in successful",
		"data": fiber.Map{
			"token": token,
		},
	})
}
