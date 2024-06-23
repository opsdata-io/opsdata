package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/opsdata-io/opsdata/pkg/models"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

type CreateUploadLinkRequest struct {
	Customer   string `json:"customer" binding:"required"`
	CaseNumber string `json:"caseNumber" binding:"required"`
	Subject    string `json:"subject" binding:"required"`
	Notes      string `json:"notes" binding:"required"`
}

func CreateUploadLink(c *fiber.Ctx) error {
	var req CreateUploadLinkRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	link := models.UploadLink{
		Customer:   req.Customer,
		CaseNumber: req.CaseNumber,
		Subject:    req.Subject,
		Notes:      req.Notes,
		// Assign other fields as needed
	}

	if err := utils.SaveUploadLink(&link); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save upload link"})
	}

	return c.Status(http.StatusCreated).JSON(link)
}

func UploadFile(c *fiber.Ctx) error {
	uploadLinkID := c.Params("link")

	// Retrieve the upload link from the database
	uploadLink, err := utils.GetUploadLink(uploadLinkID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Upload link not found"})
	}

	// Parse the file from the form data
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse file"})
	}

	// Open the file
	fileContent, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot open file"})
	}
	defer fileContent.Close()

	// Read the file content
	fileData, err := ioutil.ReadAll(fileContent)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot read file"})
	}

	// Encrypt the file
	encryptedData, err := utils.EncryptFile(fileData, []byte(utils.GetEncryptionKey()))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot encrypt file"})
	}

	// Save the file to S3
	fileID := uuid.New().String()
	if err := utils.UploadToS3(encryptedData, fileID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot upload file"})
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot save file metadata"})
	}

	return c.JSON(fiber.Map{"message": "File uploaded successfully"})
}
