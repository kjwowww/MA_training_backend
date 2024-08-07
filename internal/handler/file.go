package handler

import (
	"context"
	"ma-backend-training/internal/service"
	"net/http"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

// UploadFile handler
func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"status":  false,
			"message": "file upload error",
			"data":    "",
		})
	}

	filename := filepath.Join("static", file.Filename)
	if err := c.SaveFile(file, filename); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"status":  false,
			"message": "file save error",
			"data":    "",
		})
	}

	err = service.SaveFileInfo(context.Background(), file.Filename, file.Size)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"status":  false,
			"message": "file info save error",
			"data":    "",
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"status":  true,
		"message": "file upload success",
		"data":    file.Filename,
	})
}

// GetFiles handler
func GetFiles(c *fiber.Ctx) error {
	files, err := service.GetFiles(context.Background())
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
		"message": "get files success",
		"data":    files,
	})
}

// DeleteFile handler
func DeleteFile(c *fiber.Ctx) error {
	fileID := c.Params("file_id")

	err := service.DeleteFile(context.Background(), fileID)
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
		"message": "delete file success",
		"data":    "",
	})
}
