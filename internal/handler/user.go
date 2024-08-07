package handler

import (
	"context"
	"ma-backend-training/internal/model"
	"ma-backend-training/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetUsers handler
func GetUsers(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	rows := c.QueryInt("row", 10)
	keyword := c.Query("keyword", "")

	users, err := service.GetUsers(context.Background(), page, rows, keyword)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"status":  false,
			"message": err.Error(),
			"data":    "",
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"status":  true,
		"message": "get users success",
		"data":    users,
	})
}

// DeleteUser handler
func DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	err := service.DeleteUser(context.Background(), userID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"status":  false,
			"message": err.Error(),
			"data":    "",
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"status":  true,
		"message": "delete user success",
		"data":    "",
	})
}

// UpdateUser handler
func UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	type request struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"status":  false,
			"message": "invalid request",
			"data":    "",
		})
	}

	err := service.UpdateUser(context.Background(), userID, req.FirstName, req.LastName, req.Username)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"status":  false,
			"message": err.Error(),
			"data":    "",
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"status":  true,
		"message": "update user success",
		"data":    req,
	})
}

// ResetPassword handler
func ResetPassword(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	var req model.ResetPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"status":  false,
			"message": "invalid request",
			"data":    "",
		})
	}

	err := service.ResetUserPassword(context.Background(), userID, req.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"status":  false,
			"message": err.Error(),
			"data":    "",
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"status":  true,
		"message": "reset password success",
		"data":    "",
	})
}
