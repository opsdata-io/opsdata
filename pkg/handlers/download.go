package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/models"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

func DownloadFiles(c *fiber.Ctx) error {
	// Extract user from context
	user := c.Locals("user").(*models.User)

	// Retrieve files uploaded by the user
	files, err := utils.GetFilesByUserID(user.ID)
	if err != nil {
		log.Println("Error retrieving files:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve files"})
	}

	fileList := make([]fiber.Map, len(files))
	for i, file := range files {
		downloadLink, err := utils.GenerateDownloadLink(file.ID)
		if err != nil {
			log.Println("Error generating download link:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate download link"})
		}
		fileList[i] = fiber.Map{
			"id":           file.ID,
			"fileName":     file.FileName,
			"downloadLink": downloadLink,
		}
	}

	return c.JSON(fileList)
}
