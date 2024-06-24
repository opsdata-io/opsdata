package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/models"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

// DownloadFiles handles downloading files uploaded by the authenticated user
// @Summary Download files uploaded by the user
// @Description Retrieves files uploaded by the authenticated user with download links
// @Tags Files
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} map[string]interface{} "Successful operation"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /download [get]
func DownloadFiles(c *fiber.Ctx) error {
	// Extract user information from context
	user := c.Locals("user").(*models.User)

	// Retrieve files uploaded by the user from the database
	files, err := utils.GetFilesByUserID(user.ID)
	if err != nil {
		log.Println("Error retrieving files:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Could not retrieve files"})
	}

	// Prepare a list to store file information including download links
	fileList := make([]map[string]interface{}, len(files))
	for i, file := range files {
		// Generate a download link for each file
		downloadLink, err := utils.GenerateDownloadLink(file.ID)
		if err != nil {
			log.Println("Error generating download link:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Could not generate download link"})
		}
		// Populate file information into the list
		fileList[i] = map[string]interface{}{
			"id":           file.ID,
			"fileName":     file.FileName,
			"downloadLink": downloadLink,
		}
	}

	// Return the list of files with download links as JSON response
	return c.JSON(fileList)
}
