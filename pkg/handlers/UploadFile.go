package handlers

import (
	"io"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/opsdata-io/opsdata/pkg/models"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

// UploadFile handles uploading a file to the specified upload link
// @Summary Upload a file
// @Description Uploads a file to the specified upload link
// @Tags Uploads
// @Accept multipart/form-data
// @Produce json
// @Param link path string true "Upload link ID"
// @Param file formData file true "File to upload"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{} "File uploaded successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Upload link not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/upload/{link} [post]
func UploadFile(c *fiber.Ctx) error {
	uploadLinkID := c.Params("link")

	// Retrieve the upload link from the database
	uploadLink, err := utils.GetUploadLink(uploadLinkID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{"error": "Upload link not found"})
	}

	// Parse the file from the form data
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Cannot parse file"})
	}

	// Open the file
	fileContent, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Cannot open file"})
	}
	defer fileContent.Close()

	// Read the file content
	fileData, err := io.ReadAll(fileContent)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Cannot read file"})
	}

	// Encrypt the file content
	encryptedData, err := utils.EncryptFile(fileData, []byte(utils.GetEncryptionKey()))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Cannot encrypt file"})
	}

	// Generate a unique ID for the file
	fileID := uuid.New().String()

	// Upload the encrypted file to S3
	if err := utils.UploadToS3(encryptedData, fileID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Cannot upload file"})
	}

	// Save file metadata to the database
	fileMetadata := &models.FileMetadata{
		ID:        fileID,
		LinkID:    uploadLink.ID,
		FileName:  file.Filename,
		CreatedAt: time.Now(),
	}
	if err := utils.SaveFileMetadata(fileMetadata); err != nil {
		log.Println("Error saving file metadata:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Cannot save file metadata"})
	}

	return c.JSON(map[string]interface{}{"message": "File uploaded successfully"})
}
