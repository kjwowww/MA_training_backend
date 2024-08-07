package handler

import (
	"ma-backend-training/internal/service"

	"github.com/gofiber/fiber/v2"
)

// GetAPIHandler handles calling an external GET API
func GetAPIHandler(c *fiber.Ctx) error {
	postID := c.Params("post_id")
	if postID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "post_id is required",
		})
	}

	result, err := service.CallingGetAPIService(postID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// PostAPIHandler handles calling an external POST API
func PostAPIHandler(c *fiber.Ctx) error {
	url := "http://api.example.com/posts" // Replace with actual URL

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	result, err := service.CallingPostAPIService(url, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
